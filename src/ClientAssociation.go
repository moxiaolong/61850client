package src

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

	initiateResponseMmsPdu := NewMMSpdu()

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
	if responsePdu.initiateErrorPDU != nil {
		Throw("Got response error of class: ", responsePdu.initiateErrorPDU.ErrorClass)
	}

	if responsePdu.initiateResponsePDU == nil {
		c.acseAssociation.disconnect()
		Throw("Error decoding InitiateResponse Pdu")
	}

	initiateResponsePDU := responsePdu.initiateResponsePDU

	if initiateResponsePDU.LocalDetailCalled != nil {
		c.negotiatedMaxPduSize = initiateResponsePDU.LocalDetailCalled.intValue()
	}

	negotiatedMaxServOutstandingCalling :=
		initiateResponsePDU.NegotiatedMaxServOutstandingCalling.intValue()

	negotiatedMaxServOutstandingCalled :=
		initiateResponsePDU.NegotiatedMaxServOutstandingCalled.intValue()

	var negotiatedDataStructureNestingLevel int
	if initiateResponsePDU.NegotiatedDataStructureNestingLevel != nil {
		negotiatedDataStructureNestingLevel =
			initiateResponsePDU.NegotiatedDataStructureNestingLevel.intValue()
	} else {
		negotiatedDataStructureNestingLevel = proposedDataStructureNestingLevel
	}

	if c.negotiatedMaxPduSize < 64 || c.negotiatedMaxPduSize > proposedMaxPduSize || negotiatedMaxServOutstandingCalling > proposedMaxServOutstandingCalling || negotiatedMaxServOutstandingCalling < 0 || negotiatedMaxServOutstandingCalled > proposedMaxServOutstandingCalled || negotiatedMaxServOutstandingCalled < 0 || negotiatedDataStructureNestingLevel > proposedDataStructureNestingLevel || negotiatedDataStructureNestingLevel < 0 {

		c.acseAssociation.disconnect()
		Throw("Error negotiating parameters")
	}

	version :=
		initiateResponsePDU.InitResponseDetail.NegotiatedVersionNumber.intValue()
	if version != 1 {
		Throw("Unsupported version number was negotiated.")
	}

	c.servicesSupported = initiateResponsePDU.InitResponseDetail.ServicesSupportedCalled.value
	if (c.servicesSupported[0] & 0x40) != 0x40 {
		Throw("Obligatory services are not supported by the server.")
	}
}

func (a *ClientAssociation) Close() {

}

func (a *ClientAssociation) RetrieveModel() ServerModel {
	return ServerModel{}
}
