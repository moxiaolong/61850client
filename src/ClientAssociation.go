package src

import "C"

type ClientAssociation struct {
	ServerModel          *ServerModel
	responseTimeout      int
	negotiatedMaxPduSize int
	reportListener       *EventListener
	acseAssociation      *AcseAssociation
	clientReceiver       *ClientReceiver
	servicesSupported    []byte
}

func NewClientAssociation(address string, port int, acseSap *ClientAcseSap, proposedMaxPduSize int,
	proposedMaxServOutstandingCalling int, proposedMaxServOutstandingCalled int, proposedDataStructureNestingLevel int,
	servicesSupportedCalling []byte, responseTimeout int, messageFragmentTimeout int, reportListener *EventListener) *ClientAssociation {

	c := &ClientAssociation{}
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

	reverseOStream := NewReverseByteArrayOutputStream(500, true)
	initiateRequestMMSpdu.encode(reverseOStream)

	c.acseAssociation =
		acseSap.associate(
			address,
			port,
			reverseOStream.getByteBuffer())

	initResponse := c.acseAssociation.getAssociateResponseAPdu()

	initiateResponseMmsPdu := newMMSpdu()

	initiateResponseMmsPdu.decode(NewByteBufferInputStream(initResponse))

	c.handleInitiateResponse(
		initiateResponseMmsPdu,
		proposedMaxPduSize,
		proposedMaxServOutstandingCalling,
		proposedMaxServOutstandingCalled,
		proposedDataStructureNestingLevel)

	c.acseAssociation.MessageTimeout = 0
	c.clientReceiver = NewClientReceiver(c.negotiatedMaxPduSize)
	c.clientReceiver.start()
	return c
}

func (c *ClientAssociation) handleInitiateResponse(responsePdu *MMSpdu, proposedMaxPduSize int, proposedMaxServOutstandingCalling int, proposedMaxServOutstandingCalled int, proposedDataStructureNestingLevel int) {
	if responsePdu.InitiateErrorPDU != nil {
		Throw("Got response error of class: ", responsePdu.InitiateErrorPDU.ErrorClass)
	}

	if responsePdu.InitiateResponsePDU == nil {
		c.acseAssociation.disconnect()
		Throw("Error decoding InitiateResponse Pdu")
	}

	initiateResponsePdu := responsePdu.InitiateResponsePDU

	if initiateResponsePdu.LocalDetailCalled != nil {
		c.negotiatedMaxPduSize = initiateResponsePdu.LocalDetailCalled.intValue()
	}

	negotiatedMaxServOutstandingCalling :=
		initiateResponsePdu.NegotiatedMaxServOutstandingCalling.intValue()

	negotiatedMaxServOutstandingCalled :=
		initiateResponsePdu.NegotiatedMaxServOutstandingCalled.intValue()

	var negotiatedDataStructureNestingLevel int
	if initiateResponsePdu.NegotiatedDataStructureNestingLevel != nil {
		negotiatedDataStructureNestingLevel =
			initiateResponsePdu.NegotiatedDataStructureNestingLevel.intValue()
	} else {
		negotiatedDataStructureNestingLevel = proposedDataStructureNestingLevel
	}

	if c.negotiatedMaxPduSize < 64 || c.negotiatedMaxPduSize > proposedMaxPduSize || negotiatedMaxServOutstandingCalling > proposedMaxServOutstandingCalling || negotiatedMaxServOutstandingCalling < 0 || negotiatedMaxServOutstandingCalled > proposedMaxServOutstandingCalled || negotiatedMaxServOutstandingCalled < 0 || negotiatedDataStructureNestingLevel > proposedDataStructureNestingLevel || negotiatedDataStructureNestingLevel < 0 {

		c.acseAssociation.disconnect()
		Throw("Error negotiating parameters")
	}

	version :=
		initiateResponsePdu.InitResponseDetail.NegotiatedVersionNumber.intValue()
	if version != 1 {
		Throw("Unsupported version number was negotiated.")
	}

	c.servicesSupported = initiateResponsePdu.InitResponseDetail.ServicesSupportedCalled.value
	if (c.servicesSupported[0] & 0x40) != 0x40 {
		Throw("Obligatory services are not supported by the server.")
	}
}

func (a *ClientAssociation) close() {

}

func (a *ClientAssociation) retrieveModel() ServerModel {
	return ServerModel{}
}
