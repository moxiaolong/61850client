package src

import (
	"bytes"
	"container/list"
	"sync"
)

type ClientReceiver struct {
	pduBuffer             *bytes.Buffer
	maxMmsPduSize         int
	lock                  *sync.Mutex
	closed                bool
	acseAssociation       *AcseAssociation
	reportListener        *ClientEventListener
	incomingResponses     *list.List
	incomingResponsesLock *sync.Mutex
	expectedResponseId    int
	association           *ClientAssociation
}

func NewClientReceiver(maxMmsPduSize int, association *ClientAssociation) *ClientReceiver {
	return &ClientReceiver{maxMmsPduSize: maxMmsPduSize, closed: false, incomingResponses: list.New(), expectedResponseId: -1, association: association}
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
		buffer = r.acseAssociation.receive(r.pduBuffer)
		decodedResponsePdu := NewMMSpdu()
		decodedResponsePdu.decode(bytes.NewBuffer(buffer))

		if decodedResponsePdu.unconfirmedPDU != nil {
			if decodedResponsePdu.unconfirmedPDU.Service.InformationReport.VariableAccessSpecification.ListOfVariable != nil {
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
			r.incomingResponsesLock.Lock()
			{
				if r.expectedResponseId == -1 {
					// Discarding Reject MMS PDU because no listener for request was found.
					continue
				} else if decodedResponsePdu.rejectPDU.OriginalInvokeID.value != r.expectedResponseId {
					// Discarding Reject MMS PDU because no listener with fitting invokeID was found.
					continue
				} else {
					r.incomingResponses.PushBack(decodedResponsePdu)
				}
			}
			r.incomingResponsesLock.Unlock()
		} else if decodedResponsePdu.confirmedErrorPDU != nil {
			r.incomingResponsesLock.Lock()

			if r.expectedResponseId == -1 {
				// Discarding ConfirmedError MMS PDU because no listener for request was found.
				continue
			} else if decodedResponsePdu.confirmedErrorPDU.invokeID.value != r.expectedResponseId {
				// Discarding ConfirmedError MMS PDU because no listener with fitting invokeID was
				// found.
				continue
			} else {
				r.incomingResponses.PushBack(decodedResponsePdu)
			}
			r.incomingResponsesLock.Unlock()
		} else {
			r.incomingResponsesLock.Lock()

			if r.expectedResponseId == -1 {
				// Discarding ConfirmedResponse MMS PDU because no listener for request was found.
				continue
			} else if decodedResponsePdu.confirmedResponsePDU.invokeID.value != r.expectedResponseId {
				// Discarding ConfirmedResponse MMS PDU because no listener with fitting invokeID
				// was
				// found.
				continue
			} else {
				r.incomingResponses.PushBack(decodedResponsePdu)
			}
			r.incomingResponsesLock.Unlock()
		}

	}

}

func (r *ClientReceiver) close(err any) {
	r.lock.Lock()
	if r.closed == false {
		r.closed = true
		r.acseAssociation.disconnect()

		if r.reportListener != nil {
			go r.reportListener.associationClosed(err)
		}

		mmsPdu := NewMMSpdu()
		mmsPdu.confirmedRequestPDU = NewConfirmedRequestPDU()

		r.incomingResponses.PushBack(mmsPdu)

	}

	r.lock.Unlock()
}

func (r *ClientReceiver) processReport(mmsPdu *MMSpdu) { //*Report
	if mmsPdu.unconfirmedPDU == nil {
		Throw("getReport: Error decoding server response")
	}

	unconfirmedRes := mmsPdu.unconfirmedPDU

	if unconfirmedRes.Service == nil {
		Throw("getReport: Error decoding server response")
	}

	unconfirmedServ := unconfirmedRes.Service

	if unconfirmedServ.InformationReport == nil {
		Throw("getReport: Error decoding server response")
	}

	listRes :=
		unconfirmedServ.InformationReport.listOfAccessResult.seqOf

	index := 0

	if listRes[index].Success.visibleString == nil {
		Throw("processReport: report does not contain RptID")
	}
	index++
	rptId := listRes[index].Success.visibleString.toString()

	if listRes[index].Success.bitString == nil {
		Throw("processReport: report does not contain OptFlds")
	}

	optFlds := NewBdaOptFlds(NewObjectReference("none"), nil)
	index++
	optFlds.value = listRes[(index)].Success.bitString.value

	sqNum := -1
	if optFlds.isSequenceNumber() {
		index++
		sqNum = listRes[index].Success.Unsigned.value
	}

	timeOfEntry := -1
	if optFlds.isReportTimestamp() {
		timeOfEntry := NewBdaEntryTime(NewObjectReference("none"), nil, "", false, false)
		index++
		timeOfEntry.setValueFromMmsDataObj(listRes[index].Success)
	}

	dataSetRef := ""
	if optFlds.isDataSetName() {
		index++
		dataSetRef = listRes[index].Success.visibleString.toString()
	} else {
		urcbs := r.association.ServerModel.urcbs
		for s := range urcbs {
			urcb := urcbs[s]
			if urcb.getRptId() != nil && urcb.getRptId().getStringValue() == (rptId) || urcb.objectReference.toString() == (rptId) {
				dataSetRef = urcb.getDatSet().getStringValue()
				break
			}
		}

		if dataSetRef == null {
			for Brcb
		brcb:
			serverModel.getBrcbs()) {
				if brcb.getRptId() != null && brcb.getRptId().getStringValue().equals(rptId)
|| brcb.getReference().toString().equals(rptId)) {
dataSetRef = brcb.getDatSet().getStringValue();
break;
}
}
}
}
if (dataSetRef == null) {
throw new ServiceError(
ServiceError.FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT,
"unable to find RCB that matches the given RptID in the report.");
}
dataSetRef = dataSetRef.replace('$', '.')

DataSet dataSet = serverModel.getDataSet(dataSetRef)
if (dataSet == null) {
throw new ServiceError(
ServiceError.FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT,
"unable to find data set that matches the given data set reference of the report.");
}

Boolean bufOvfl = null
if (optFlds.isBufferOverflow()) {
bufOvfl = listRes.get(index++).getSuccess().getBool().value;
}

BdaOctetString entryId = null
if (optFlds.isEntryId()) {
entryId = new BdaOctetString(new ObjectReference("none"), null, "", 8, false, false);
entryId.setValue(listRes.get(index++).getSuccess().getOctetString().value);
}

Long confRev = null
if (optFlds.isConfigRevision()) {
confRev = listRes.get(index++).getSuccess().getUnsigned().longValue();
}

Integer subSqNum = null
boolean moreSegmentsFollow = false
if (optFlds.isSegmentation()) {
subSqNum = listRes.get(index++).getSuccess().getUnsigned().intValue();
moreSegmentsFollow = listRes.get(index++).getSuccess().getBool().value;
}

boolean[] inclusionBitString =
listRes.get(index++).getSuccess().getBitString().getValueAsBooleans()
int numMembersReported = 0
for (boolean bit: inclusionBitString) {
if (bit) {
numMembersReported++;
}
}

if (optFlds.isDataReference()) {
// this is just to move the index to the right place
// The next part will process the changes to the values
// without the dataRefs
index += numMembersReported;
}

List<FcModelNode> reportedDataSetMembers = new ArrayList<>(numMembersReported)
int dataSetIndex = 0
for (FcModelNode dataSetMember: dataSet.getMembers()) {
if (inclusionBitString[dataSetIndex]) {
AccessResult accessRes = listRes.get(index++);
FcModelNode dataSetMemberCopy = (FcModelNode) dataSetMember.copy();
dataSetMemberCopy.setValueFromMmsDataObj(accessRes.getSuccess());
reportedDataSetMembers.add(dataSetMemberCopy);
}
dataSetIndex++;
}

List<BdaReasonForInclusion> reasonCodes = null
if (optFlds.isReasonForInclusion()) {
reasonCodes = new ArrayList<>(dataSet.getMembers().size());
for (int i = 0; i < dataSet.getMembers().size(); i++) {
if (inclusionBitString[i]) {
BdaReasonForInclusion reasonForInclusion = new BdaReasonForInclusion(null);
reasonCodes.add(reasonForInclusion);
byte[] reason = listRes.get(index++).getSuccess().getBitString().value;
reasonForInclusion.setValue(reason);
}
}
}
return nil
return new Report(
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
