package src

import (
	"container/list"
	"sync"
	"unsafe"
)

type ClientAssociation struct {
	ServerModel          *ServerModel
	responseTimeout      int
	negotiatedMaxPduSize int
	reportListener       *EventListener
	acseAssociation      *AcseAssociation
	clientReceiver       *ClientReceiver
	servicesSupported    []byte
	lock                 sync.Mutex
	closed               bool
	incomingResponses    *list.List
}

func NewClientAssociation(address string, port int, acseSap *ClientAcseSap, proposedMaxPduSize int,
	proposedMaxServOutstandingCalling int, proposedMaxServOutstandingCalled int, proposedDataStructureNestingLevel int,
	servicesSupportedCalling []byte, responseTimeout int, messageFragmentTimeout int, reportListener *EventListener) *ClientAssociation {

	c := &ClientAssociation{}
	c.incomingResponses = list.New()
	c.closed = false
	c.responseTimeout = responseTimeout
	acseSap.tSap.MessageFragmentTimeout = messageFragmentTimeout
	acseSap.tSap.MessageTimeout = responseTimeout
	c.negotiatedMaxPduSize = proposedMaxPduSize
	c.reportListener = reportListener

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
		c.incomingResponses.PushBack(mmsPdu)
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
	confirmedServiceResponse := encodeWriteReadDecode(serviceRequest)
	return decodeGetDataDefinitionResponse(confirmedServiceResponse, lnRef)
}

func decodeGetDataDefinitionResponse(response *ConfirmedServiceResponse, ref *ObjectReference) *LogicalNode {
	//TODO
	return nil
}

func encodeWriteReadDecode(request *ConfirmedServiceRequest) *ConfirmedServiceResponse {

	return nil
}

func (c *ClientAssociation) retrieveLogicalNodeNames(s string) []string {
	return nil
}

func (c *ClientAssociation) constructGetServerDirectoryRequest() *ConfirmedServiceRequest {
	return nil
}

func (c *ClientAssociation) encodeWriteReadDecode(request *ConfirmedServiceRequest) *ConfirmedServiceResponse {
	return nil
}

func (c *ClientAssociation) decodeGetServerDirectoryResponse(response *ConfirmedServiceResponse) []string {
	return nil
}

func (c *ClientAssociation) constructGetDirectoryRequest(name interface{}, s string, b bool) *ConfirmedServiceRequest {
	return nil
}

func (c *ClientAssociation) decodeAndRetrieveDsNamesAndDefinitions(response *ConfirmedServiceResponse, l *LogicalDevice) {

}

func (c *ClientAssociation) constructGetDataDefinitionRequest(ref *ObjectReference) *ConfirmedServiceRequest {
	return nil
}
