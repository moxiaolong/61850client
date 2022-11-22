package src

import (
	"bytes"
	"strings"
	"sync"
)

type ClientReceiver struct {
	pduBuffer          *bytes.Buffer
	maxMmsPduSize      int
	lock               *sync.Mutex
	closed             bool
	reportListener     *ClientEventListener
	expectedResponseId int
	association        *ClientAssociation
	lastIOException    string
}

func NewClientReceiver(maxMmsPduSize int, association *ClientAssociation) *ClientReceiver {
	return &ClientReceiver{maxMmsPduSize: maxMmsPduSize, closed: false, expectedResponseId: -1, association: association, pduBuffer: bytes.NewBuffer(make([]byte, maxMmsPduSize+400)),
		lock: &sync.Mutex{}}
}
func (r *ClientReceiver) start() {
	go r.run()
}

func (r *ClientReceiver) run() {
	defer func() {
		err := recover()
		if err != nil {
			r.close(err)
		}
	}()

	for {
		r.pduBuffer.Reset()
		var buffer []byte
		buffer = r.association.acseAssociation.receive(r.pduBuffer)
		decodedResponsePdu := NewMMSpdu()
		decodedResponsePdu.decode(bytes.NewBuffer(buffer))

		if decodedResponsePdu.unconfirmedPDU != nil {
			if decodedResponsePdu.unconfirmedPDU.service.informationReport.variableAccessSpecification.listOfVariable != nil {
				// Discarding LastApplError Report
			} else {
				if r.reportListener != nil {

					report := r.processReport(decodedResponsePdu)
					go func() {
						r.reportListener.newReport(report)
					}()
				} else {
					// discarding report because no ReportListener was registered.
				}
			}
		} else if decodedResponsePdu.rejectPDU != nil {
			r.association.incomingResponsesLock.Lock()
			{
				if r.expectedResponseId == -1 {
					// Discarding Reject MMS PDU because no listener for request was found.
					continue
				} else if decodedResponsePdu.rejectPDU.originalInvokeID.value != r.expectedResponseId {
					// Discarding Reject MMS PDU because no listener with fitting invokeID was found.
					continue
				} else {
					r.association.incomingResponses <- decodedResponsePdu
				}
			}
			r.association.incomingResponsesLock.Unlock()
		} else if decodedResponsePdu.confirmedErrorPDU != nil {
			r.association.incomingResponsesLock.Lock()

			if r.expectedResponseId == -1 {
				// Discarding ConfirmedError MMS PDU because no listener for request was found.
				continue
			} else if decodedResponsePdu.confirmedErrorPDU.invokeID.value != r.expectedResponseId {
				// Discarding ConfirmedError MMS PDU because no listener with fitting invokeID was
				// found.
				continue
			} else {
				r.association.incomingResponses <- decodedResponsePdu
			}
			r.association.incomingResponsesLock.Unlock()
		} else {
			r.association.incomingResponsesLock.Lock()

			if r.expectedResponseId == -1 {
				// Discarding ConfirmedResponse MMS PDU because no listener for request was found.
				continue
			} else if decodedResponsePdu.confirmedResponsePDU.invokeID.value != r.expectedResponseId {
				// Discarding ConfirmedResponse MMS PDU because no listener with fitting invokeID
				// was
				// found.
				continue
			} else {
				r.association.incomingResponses <- decodedResponsePdu
			}
			r.association.incomingResponsesLock.Unlock()
		}

	}

}

func (r *ClientReceiver) close(err any) {
	r.lock.Lock()
	if r.closed == false {
		r.closed = true
		r.association.acseAssociation.disconnect()

		if r.reportListener != nil {
			go r.reportListener.associationClosed(err)
		}

		mmsPdu := NewMMSpdu()
		mmsPdu.confirmedRequestPDU = NewConfirmedRequestPDU()

		r.association.incomingResponses <- mmsPdu

	}

	r.lock.Unlock()
}

func (r *ClientReceiver) processReport(mmsPdu *MMSpdu) *Report {
	if mmsPdu.unconfirmedPDU == nil {
		throw("getReport: Error decoding server response")
	}

	unconfirmedRes := mmsPdu.unconfirmedPDU

	if unconfirmedRes.service == nil {
		throw("getReport: Error decoding server response")
	}

	unconfirmedServ := unconfirmedRes.service

	if unconfirmedServ.informationReport == nil {
		throw("getReport: Error decoding server response")
	}

	listRes :=
		unconfirmedServ.informationReport.listOfAccessResult.seqOf

	index := 0

	if listRes[index].success.visibleString == nil {
		throw("processReport: report does not contain RptID")
	}
	index++
	rptId := listRes[index].success.visibleString.toString()

	if listRes[index].success.bitString == nil {
		throw("processReport: report does not contain OptFlds")
	}

	optFlds := NewBdaOptFlds(NewObjectReference("none"), "")
	index++
	optFlds.value = listRes[(index)].success.bitString.value

	var sqNum *int = nil
	if optFlds.isSequenceNumber() {
		index++
		sqNum = &listRes[index].success.Unsigned.value
	}

	var timeOfEntry *BdaEntryTime = nil
	if optFlds.isReportTimestamp() {
		timeOfEntry = NewBdaEntryTime(NewObjectReference("none"), "", "", false, false)
		index++
		timeOfEntry.setValueFromMmsDataObj(listRes[index].success)
	}

	dataSetRef := ""
	if optFlds.isDataSetName() {
		index++
		dataSetRef = listRes[index].success.visibleString.toString()
	} else {
		urcbs := r.association.ServerModel.urcbs
		for s := range urcbs {
			urcb := urcbs[s]
			if urcb.getRptId() != nil && urcb.getRptId().getStringValue() == (rptId) || urcb.objectReference.toString() == (rptId) {
				dataSetRef = urcb.getDatSet().getStringValue()
				break
			}
		}
	}

	if dataSetRef == "" {
		throw(
			"unable to find RCB that matches the given RptID in the report.")
	}

	dataSetRef = strings.ReplaceAll(dataSetRef, "$", ".")

	dataSet := r.association.ServerModel.getDataSet(dataSetRef)
	if dataSet == nil {
		throw(
			"unable to find data set that matches the given data set reference of the report.")
	}

	var bufOvfl *bool
	if optFlds.isBufferOverflow() {
		index++
		bufOvfl = &listRes[index].success.bool.value
	}

	var entryId *BdaOctetString = nil
	if optFlds.isEntryId() {
		entryId = NewBdaOctetString(NewObjectReference("none"), "", "", 8, false, false)
		index++
		entryId.setValue(listRes[index].success.octetString.value)
	}

	var confRev *int = nil
	if optFlds.isConfigRevision() {

		index++
		confRev = &listRes[index].success.Unsigned.value
	}

	var subSqNum *int = nil
	moreSegmentsFollow := false
	if optFlds.isSegmentation() {
		index++
		subSqNum = &listRes[index].success.Unsigned.value
		index++
		moreSegmentsFollow = listRes[index].success.bool.value
	}

	index++
	inclusionBitString := listRes[index].success.bitString.getValueAsBooleans()
	numMembersReported := 0

	for _, bit := range inclusionBitString {
		if bit {
			numMembersReported++
		}
	}

	if optFlds.isDataReference() {
		// this is just to move the index to the right place
		// The next part will process the changes to the values
		// without the dataRefs
		index += numMembersReported
	}

	reportedDataSetMembers := make([]*FcModelNode, 0)
	//reportedDataSetMembers := make([]*FcModelNode, numMembersReported)
	dataSetIndex := 0
	for _, dataSetMember := range dataSet.getMembers() {
		if inclusionBitString[dataSetIndex] {
			index++
			accessRes := listRes[index]

			//TPDO
			dataSetMemberCopy := dataSetMember.copy()
			dataSetMemberCopy.setValueFromMmsDataObj(accessRes.success)
			reportedDataSetMembers = append(reportedDataSetMembers, dataSetMemberCopy)
		}
		dataSetIndex++
	}

	var reasonCodes []*BdaReasonForInclusion = nil
	if optFlds.isReasonForInclusion() {
		//reasonCodes = make([]*BdaReasonForInclusion, len(dataSets.getMembers()))
		reasonCodes = make([]*BdaReasonForInclusion, 0)
		for i := 0; i < len(dataSet.getMembers()); i++ {
			if inclusionBitString[i] {

				reasonForInclusion := NewBdaReasonForInclusion(nil)
				reasonCodes = append(reasonCodes, reasonForInclusion)
				index++
				reason := listRes[index].success.bitString.value
				reasonForInclusion.value = reason
			}
		}
	}

	return NewReport(
		rptId,
		sqNum,
		subSqNum,
		moreSegmentsFollow,
		dataSetRef,
		bufOvfl,
		confRev,
		timeOfEntry,
		entryId,
		inclusionBitString,
		reportedDataSetMembers,
		reasonCodes)
}

func (r *ClientReceiver) removeExpectedResponse() *MMSpdu {
	r.association.incomingResponsesLock.Lock()
	r.expectedResponseId = -1
	spdu := <-r.association.incomingResponses
	r.association.incomingResponsesLock.Unlock()
	return spdu
}
