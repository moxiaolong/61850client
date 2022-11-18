package src

import "C"
import (
	"strconv"
	"sync"
	"time"
	"unsafe"
)

type ClientAssociation struct {
	ServerModel           *ServerModel
	responseTimeout       int
	negotiatedMaxPduSize  int
	reportListener        *EventListener
	acseAssociation       *AcseAssociation
	clientReceiver        *ClientReceiver
	servicesSupported     []byte
	lock                  *sync.Mutex
	closed                bool
	incomingResponses     chan *MMSpdu
	incomingResponsesLock *sync.Mutex
	invokeId              int
	reverseOStream        *ReverseByteArrayOutputStream
}

func NewClientAssociation(address string, port int, acseSap *ClientAcseSap, proposedMaxPduSize int,
	proposedMaxServOutstandingCalling int, proposedMaxServOutstandingCalled int, proposedDataStructureNestingLevel int,
	servicesSupportedCalling []byte, responseTimeout int, messageFragmentTimeout int, reportListener *EventListener) *ClientAssociation {

	c := &ClientAssociation{}
	c.lock = &sync.Mutex{}
	c.incomingResponses = make(chan *MMSpdu)
	c.incomingResponsesLock = &sync.Mutex{}
	c.closed = false
	c.responseTimeout = responseTimeout
	acseSap.tSap.MessageFragmentTimeout = messageFragmentTimeout
	acseSap.tSap.MessageTimeout = responseTimeout
	c.negotiatedMaxPduSize = proposedMaxPduSize
	c.reportListener = reportListener
	c.reverseOStream = NewReverseByteArrayOutputStream(500)

	initiateRequestMMSpdu :=
		constructInitRequestPdu(
			proposedMaxPduSize,
			proposedMaxServOutstandingCalling,
			proposedMaxServOutstandingCalled,
			proposedDataStructureNestingLevel,
			servicesSupportedCalling)

	reverseOStream := NewReverseByteArrayOutputStream(500)
	initiateRequestMMSpdu.encode(reverseOStream)

	c.acseAssociation =
		acseSap.associate(
			address,
			port,
			reverseOStream.getByteBuffer())

	initResponse := c.acseAssociation.getAssociateResponseAPdu()

	initiateResponseMmsPdu := NewMMSpdu()

	initiateResponseMmsPdu.decode(initResponse)

	c.handleInitiateResponse(
		initiateResponseMmsPdu,
		proposedMaxPduSize,
		proposedMaxServOutstandingCalling,
		proposedMaxServOutstandingCalled,
		proposedDataStructureNestingLevel)

	c.acseAssociation.MessageTimeout = 0
	c.clientReceiver = NewClientReceiver(c.negotiatedMaxPduSize, c)
	c.clientReceiver.start()
	return c
}

func (c *ClientAssociation) handleInitiateResponse(responsePdu *MMSpdu, proposedMaxPduSize int, proposedMaxServOutstandingCalling int, proposedMaxServOutstandingCalled int, proposedDataStructureNestingLevel int) {
	if responsePdu.initiateErrorPDU != nil {
		throw("Got response error of class: ") //responsePdu.initiateErrorPDU.errorClass) TODO
	}

	if responsePdu.initiateResponsePDU == nil {
		c.acseAssociation.disconnect()
		throw("Error decoding InitiateResponse Pdu")
	}

	initiateResponsePDU := responsePdu.initiateResponsePDU

	if initiateResponsePDU.localDetailCalled != nil {
		c.negotiatedMaxPduSize = initiateResponsePDU.localDetailCalled.intValue()
	}

	negotiatedMaxServOutstandingCalling :=
		initiateResponsePDU.negotiatedMaxServOutstandingCalling.intValue()

	negotiatedMaxServOutstandingCalled :=
		initiateResponsePDU.negotiatedMaxServOutstandingCalled.intValue()

	var negotiatedDataStructureNestingLevel int
	if initiateResponsePDU.negotiatedDataStructureNestingLevel != nil {
		negotiatedDataStructureNestingLevel =
			initiateResponsePDU.negotiatedDataStructureNestingLevel.intValue()
	} else {
		negotiatedDataStructureNestingLevel = proposedDataStructureNestingLevel
	}

	if c.negotiatedMaxPduSize < 64 || c.negotiatedMaxPduSize > proposedMaxPduSize || negotiatedMaxServOutstandingCalling > proposedMaxServOutstandingCalling || negotiatedMaxServOutstandingCalling < 0 || negotiatedMaxServOutstandingCalled > proposedMaxServOutstandingCalled || negotiatedMaxServOutstandingCalled < 0 || negotiatedDataStructureNestingLevel > proposedDataStructureNestingLevel || negotiatedDataStructureNestingLevel < 0 {

		c.acseAssociation.disconnect()
		throw("Error negotiating parameters")
	}

	version :=
		initiateResponsePDU.initResponseDetail.negotiatedVersionNumber.intValue()
	if version != 1 {
		throw("Unsupported version number was negotiated.")
	}

	c.servicesSupported = initiateResponsePDU.initResponseDetail.servicesSupportedCalled.value
	if (c.servicesSupported[0] & 0x40) != 0x40 {
		throw("Obligatory services are not supported by the server.")
	}
}

func (c *ClientAssociation) Close() {
	c.lock.Lock()
	if c.closed == false {
		c.closed = true
		c.acseAssociation.disconnect()
		go c.reportListener.associationClosed()

		mmsPdu := NewMMSpdu()
		mmsPdu.confirmedRequestPDU = NewConfirmedRequestPDU()
		c.incomingResponses <- mmsPdu
	}
	c.lock.Unlock()

}

func (c *ClientAssociation) RetrieveModel() *ServerModel {
	ldNames := c.retrieveLogicalDevices()
	lnNames := make([][]string, len(ldNames))

	for i := 0; i < len(ldNames); i++ {
		lnNames = append(lnNames, c.retrieveLogicalNodeNames(ldNames[i]))
	}
	lds := make([]*LogicalDevice, 0)
	for i := 0; i < len(ldNames); i++ {
		lns := make([]*LogicalNode, 0)
		for j := 0; j < len(lnNames[i]); j++ {
			lns = append(lns, c.retrieveDataDefinitions(
				NewObjectReference(ldNames[i]+"/"+lnNames[i][j])))

		}
		lds = append(lds, NewLogicalDevice(NewObjectReference(ldNames[i]), lns))
	}

	c.ServerModel = NewServerModel(lds, nil)

	c.updateDataSets()

	return c.ServerModel

}

func (c *ClientAssociation) retrieveLogicalDevices() []string {
	serviceRequest := c.constructGetServerDirectoryRequest()
	confirmedServiceResponse := c.encodeWriteReadDecode(serviceRequest)
	return c.decodeGetServerDirectoryResponse(confirmedServiceResponse)
}

func (c *ClientAssociation) updateDataSets() {
	if c.ServerModel == nil {
		throw("Before calling this function you have to get the ServerModel using the retrieveModel() function")
	}
	lds := c.ServerModel.children
	for _, ld := range lds {
		serviceRequest :=
			c.constructGetDirectoryRequest(ld.objectReference.getName(), "", false)
		confirmedServiceResponse := c.encodeWriteReadDecode(serviceRequest)
		pointer := unsafe.Pointer(ld)
		c.decodeAndRetrieveDsNamesAndDefinitions(confirmedServiceResponse, (*LogicalDevice)(pointer))
	}

}

func (c *ClientAssociation) retrieveDataDefinitions(lnRef *ObjectReference) *LogicalNode {
	serviceRequest := c.constructGetDataDefinitionRequest(lnRef)
	confirmedServiceResponse := c.encodeWriteReadDecode(serviceRequest)
	return decodeGetDataDefinitionResponse(confirmedServiceResponse, lnRef)
}

func decodeGetDataDefinitionResponse(response *ConfirmedServiceResponse, ref *ObjectReference) *LogicalNode {
	//TODO
	return nil
}

func (c *ClientAssociation) encodeWriteReadDecode(serviceRequest *ConfirmedServiceRequest) *ConfirmedServiceResponse {
	currentInvokeId := c.getInvokeId()

	confirmedRequestPdu := NewConfirmedRequestPDU()
	confirmedRequestPdu.invokeID = NewUnsigned32(currentInvokeId)
	confirmedRequestPdu.service = serviceRequest

	requestPdu := NewMMSpdu()
	requestPdu.confirmedRequestPDU = confirmedRequestPdu

	c.reverseOStream.reset()

	func() {
		defer func() {
			r := recover()
			if r != nil {
				c.clientReceiver.close(r)
				panic(r)
			}
		}()
		requestPdu.encode(c.reverseOStream)
	}()

	c.clientReceiver.expectedResponseId = currentInvokeId

	func() {
		defer func() {
			r := recover()
			if r != nil {
				throw("Error sending packet.")
				c.clientReceiver.close(r)
				panic(r)
			}
		}()
		c.acseAssociation.sendByteBuffer(c.reverseOStream.getByteBuffer())
	}()

	var decodedResponsePdu *MMSpdu = nil

	func() {
		defer func() {
			r := recover()
			if r != nil {

			}
		}()
		if c.responseTimeout == 0 {
			if len(c.incomingResponses) > 0 {
				decodedResponsePdu = <-c.incomingResponses
			}
		} else {
			timeOut := time.After(time.Duration(c.responseTimeout) * time.Millisecond)
			select {
			case decodedResponsePdu = <-c.incomingResponses:
				break
			case <-timeOut:
				panic("time out")
			}
		}
	}()

	if decodedResponsePdu == nil {
		decodedResponsePdu = c.clientReceiver.removeExpectedResponse()
		if decodedResponsePdu == nil {
			throw("Service error TIMEOUT_ERROR")
		}
	}

	if decodedResponsePdu.confirmedResponsePDU != nil {
		c.incomingResponses <- decodedResponsePdu
		throw("connection was closed", c.clientReceiver.lastIOException)
	}

	testForInitiateErrorResponse(decodedResponsePdu)
	testForErrorResponse(decodedResponsePdu)
	testForRejectResponse(decodedResponsePdu)

	confirmedResponsePdu := decodedResponsePdu.confirmedResponsePDU
	if confirmedResponsePdu == nil {
		throw("Response PDU is not a confirmed response pdu")
	}

	return confirmedResponsePdu.service

}

func testForRejectResponse(mmsResponsePdu *MMSpdu) {
	if mmsResponsePdu.rejectPDU == nil {
		return
	}

	rejectReason := mmsResponsePdu.rejectPDU.rejectReason
	if rejectReason != nil {
		if rejectReason.pduError != nil {
			if rejectReason.pduError.value == 1 {
				throw(
					" PARAMETER_VALUE_INCONSISTENTMMS reject: type: \"pdu-error\", reject code: \"invalid-pdu\"")
			}
		}
	}
	throw(" UNKNOWN MMS confirmed error.")
}

func testForErrorResponse(mmsResponsePdu *MMSpdu) {
	if mmsResponsePdu.confirmedResponsePDU == nil {
		return
	}

	errClass := mmsResponsePdu.confirmedErrorPDU.serviceError.errorClass

	if errClass != nil {
		if errClass.access != nil {
			if errClass.access.value == 3 {
				throw(
					"ACCESS_VIOLATION MMS confirmed error: class: \"access\", error code: \"object-access-denied\"")
			} else if errClass.access.value == 2 {
				throw(
					" INSTANCE_NOT_AVAILABLEMMS confirmed error: class: \"access\", error code: \"object-non-existent\"")
			}

		} else if errClass.file != nil {
			if errClass.file.value == 7 {
				throw(
					"FILE_NONE_EXISTENT  MMS confirmed error: class: \"file\", error code: \"file-non-existent\"")
			}
		}
	}

	if mmsResponsePdu.confirmedErrorPDU.serviceError.additionalDescription != nil {
		throw(
			"UNKNOWN MMS confirmed error. Description: ",
			mmsResponsePdu.confirmedErrorPDU.serviceError.additionalDescription.toString())
	}
	throw("UNKNOWN  MMS confirmed error.")
}

func testForInitiateErrorResponse(mmsResponsePdu *MMSpdu) {
	if mmsResponsePdu.initiateResponsePDU != nil {

		errClass := mmsResponsePdu.initiateErrorPDU.errorClass
		if errClass != nil {
			if errClass.vmdState != nil {
				throw("FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"vmd_state\" with val: ", strconv.Itoa(errClass.vmdState.value))
			}
			if errClass.applicationReference != nil {
				throw("FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"application_reference\" with val: ", strconv.Itoa(errClass.applicationReference.value))
			}
			if errClass.definition != nil {
				throw("FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"definition\" with val: ", strconv.Itoa(errClass.definition.value))
			}
			if errClass.resource != nil {
				throw(
					" FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"resource\" with val: ", strconv.Itoa(errClass.resource.value))
			}
			if errClass.service != nil {
				throw("FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"service\" with val: ", strconv.Itoa(errClass.service.value))
			}
			if errClass.servicePreempt != nil {
				throw(

					"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT  error class \"service_preempt\" with val: " + strconv.Itoa(errClass.servicePreempt.value))
			}
			if errClass.timeResolution != nil {
				throw(

					"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"time_resolution\" with val: " + strconv.Itoa(errClass.timeResolution.value))
			}
			if errClass.access != nil {
				throw(
					"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"access\" with val: " + strconv.Itoa(errClass.access.value))
			}
			if errClass.initiate != nil {
				throw(
					"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"initiate\" with val: " + strconv.Itoa(errClass.initiate.value))
			}
			if errClass.conclude != nil {
				throw(
					"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"conclude\" with val: " + strconv.Itoa(errClass.conclude.value))
			}
			if errClass.cancel != nil {
				throw("FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"cancel\" with val: ", strconv.Itoa(errClass.cancel.value))
			}
			if errClass.file != nil {
				throw(

					"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"file\" with val: " + strconv.Itoa(errClass.file.value))
			}
			if errClass.others != nil {
				throw(
					"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT error class \"others\" with val: " + strconv.Itoa(errClass.others.value))
			}
		}
		throw(
			"FAILED_DUE_TO_COMMUNICATIONS_CONSTRAINT unknown error class")
	}
}

func (c *ClientAssociation) retrieveLogicalNodeNames(s string) []string {
	return nil
}

func (c *ClientAssociation) constructGetServerDirectoryRequest() *ConfirmedServiceRequest {
	objectClass := NewObjectClass()
	objectClass.basicObjectClass = NewBerInteger(nil, 9)

	objectScope := NewObjectScope()
	objectScope.vmdSpecific = NewBerNull()

	getNameListRequest := NewGetNameListRequest()
	getNameListRequest.objectClass = objectClass
	getNameListRequest.objectScope = objectScope

	confirmedServiceRequest := NewConfirmedServiceRequest()
	confirmedServiceRequest.getNameList = getNameListRequest

	return confirmedServiceRequest
}

func (c *ClientAssociation) decodeGetServerDirectoryResponse(response *ConfirmedServiceResponse) []string {
	return nil
}

func (c *ClientAssociation) constructGetDirectoryRequest(name interface{}, s string, b bool) *ConfirmedServiceRequest {
	return nil
}

func (c *ClientAssociation) decodeAndRetrieveDsNamesAndDefinitions(response *ConfirmedServiceResponse, l *LogicalDevice) {

}

func (c *ClientAssociation) constructGetDataDefinitionRequest(lnRef *ObjectReference) *ConfirmedServiceRequest {
	domainSpec := NewDomainSpecific()
	domainSpec.domainID = NewIdentifier([]byte(lnRef.get(0)))
	domainSpec.itemID = NewIdentifier([]byte(lnRef.get(1)))

	objectName := NewObjectName()
	objectName.domainSpecific = domainSpec

	getVariableAccessAttributesRequest := NewGetVariableAccessAttributesRequest()
	getVariableAccessAttributesRequest.name = objectName

	confirmedServiceRequest := NewConfirmedServiceRequest()
	confirmedServiceRequest.getVariableAccessAttributes = getVariableAccessAttributesRequest

	return confirmedServiceRequest
}

func (c *ClientAssociation) getInvokeId() int {
	c.invokeId = (c.invokeId + 1) % 2147483647
	return c.invokeId
}
