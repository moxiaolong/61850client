package src

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

type AcseAssociation struct {
	MessageTimeout          int
	associateResponseAPDU   *bytes.Buffer
	tConnection             *TConnection
	pSelLocalBerOctetString *RespondingPresentationSelector
	tSap                    *ClientTSap
	connected               bool
}

func (a *AcseAssociation) getAssociateResponseAPdu() *bytes.Buffer {
	returnBuffer := a.associateResponseAPDU
	a.associateResponseAPDU = nil
	return returnBuffer
}

func (a *AcseAssociation) disconnect() {
	a.connected = false
	if a.tConnection != nil {
		a.tConnection.disconnect()
	}
}

func (a *AcseAssociation) startAssociation(payload *bytes.Buffer, address string, port int, sSelRemote []byte, sSelLocal []byte, pSelRemote []byte, tSAP *ClientTSap, apTitleCalled []int, apTitleCalling []int, aeQualifierCalled int, aeQualifierCalling int) {
	if a.connected {
		Throw("io")
	}

	called_ap_title := NewAPTitle()
	called_ap_title.ApTitleForm2 = NewApTitleForm2(apTitleCalled)
	calling_ap_title := NewAPTitle()
	calling_ap_title.ApTitleForm2 = NewApTitleForm2(apTitleCalling)

	called_ae_qualifier := NewAEQualifier()
	called_ae_qualifier.AeQualifierForm2 = NewAEQualifierForm2(aeQualifierCalled)
	calling_ae_qualifier := NewAEQualifier()
	calling_ae_qualifier.AeQualifierForm2 = NewAEQualifierForm2(aeQualifierCalling)

	encoding := NewMyexternalEncoding()
	encoding.SingleASN1Type = NewBerAny(payload.Bytes())

	myExternal := NewMyexternal()
	myExternal.DirectReference = NewBerObjectIdentifier([]byte{0x02, 0x51, 0x01}) //static
	myExternal.IndirectReference = NewBerInteger([]byte{0x01, 0x03}, 0)           //static
	myExternal.Encoding = encoding

	userInformation := NewAssociationInformation()
	userInformation.seqOf = append(userInformation.seqOf, myExternal)

	aarq := NewAARQApdu()
	aarq.ApplicationContextName = NewBerObjectIdentifier([]byte{0x05, 0x28, 0xca, 0x22, 0x02, 0x03}) //static
	aarq.CalledAPTitle = called_ap_title
	aarq.CalledAEQualifier = called_ae_qualifier
	aarq.CallingAPTitle = calling_ap_title
	aarq.CallingAEQualifier = calling_ae_qualifier
	aarq.UserInformation = userInformation

	acse := NewACSEApdu()
	acse.Aarq = aarq

	reverseOStream := NewReverseByteArrayOutputStream(200)
	acse.encode(reverseOStream)

	userData := getPresentationUserDataField(reverseOStream.getArray())

	normalModeParameter := NewCPTypeNormalModeParameters()
	normalModeParameter.CallingPresentationSelector = NewCallingPresentationSelector(a.pSelLocalBerOctetString.value)
	normalModeParameter.CalledPresentationSelector = NewCalledPresentationSelector(pSelRemote)
	normalModeParameter.PresentationContextDefinitionList = NewPresentationContextDefinitionList([]byte{0x23, 0x30, 0x0f, 0x02, 0x01, 0x01, 0x06, 0x04, 0x52, 0x01, 0x00, 0x01, 0x30, 0x04, 0x06, 0x02, 0x51, 0x01, 0x30, 0x10, 0x02, 0x01, 0x03, 0x06, 0x05, 0x28, 0xca, 0x22, 0x02, 0x01, 0x30, 0x04, 0x06, 0x02, 0x51, 0x01})
	normalModeParameter.UserData = userData

	cpType := NewCPType()
	cpType.ModeSelector = NewModeSelector()
	cpType.NormalModeParameters = normalModeParameter

	reverseOStream.reset()
	cpType.encode(reverseOStream, true)

	ssduList := make([][]byte, 1)
	ssduOffsets := make([]int, 1)
	ssduLengths := make([]int, 1)

	ssduList = append(ssduList, reverseOStream.buffer)
	ssduOffsets = append(ssduOffsets, reverseOStream.index+1)
	ssduLengths = append(ssduLengths, len(reverseOStream.buffer)-(reverseOStream.index+1))

	res :=
		a.startSConnection(
			ssduList,
			ssduOffsets,
			ssduLengths,
			address,
			port,
			tSAP,
			sSelRemote,
			sSelLocal)

	a.associateResponseAPDU = decodePConResponse(res)
}

func decodePConResponse(ppdu *bytes.Buffer) *bytes.Buffer {
	cpa_ppdu := NewCPAPPDU()
	cpa_ppdu.decode(ppdu)

	acseApdu := NewACSEApdu()
	acseApdu.decode(ppdu)
	return bytes.NewBuffer(acseApdu.Aare.UserInformation.Myexternal[0].Encoding.SingleASN1Type.value)
}

func (a *AcseAssociation) startSConnection(ssduList [][]byte, ssduOffsets []int, ssduLengths []int, address string, port int, tSAP *ClientTSap, sSelRemote []byte, sSelLocal []byte) *bytes.Buffer {
	if a.connected {
		Throw("io error")
	}

	spduHeader := make([]byte, 24)
	idx := 0
	// byte[] res = null;
	ssduLength := 0
	for _, item := range ssduLengths {
		ssduLength += item
	}
	// write ISO 8327-1 Header
	// SPDU Type: CONNECT (13)
	spduHeader[idx] = 0x0d
	idx++
	// Length: length of session user data + 22 ( header data after
	// length field )
	spduHeader[idx] = (byte)((ssduLength + 22) & 0xff)
	idx++

	// -- start Connect Accept Item
	// Parameter type: Connect Accept Item (5)
	spduHeader[idx] = 0x05
	idx++
	// Parameter length
	spduHeader[idx] = 0x06
	idx++

	// Protocol options:
	// Parameter Type: Protocol Options (19)
	spduHeader[idx] = 0x13
	idx++
	// Parameter length
	spduHeader[idx] = 0x01
	idx++
	// flags: (.... ...0 = Able to receive extended concatenated SPDU:
	// False)
	spduHeader[idx] = 0x00
	idx++

	// Version number:
	// Parameter type: Version Number (22)
	spduHeader[idx] = 0x16
	idx++
	// Parameter length
	spduHeader[idx] = 0x01
	idx++
	// flags: (.... ..1. = Protocol Version 2: True)
	spduHeader[idx] = 0x02
	idx++
	// -- end Connect Accept Item

	// Session Requirement
	// Parameter type: Session Requirement (20)
	spduHeader[idx] = 0x14
	idx++
	// Parameter length
	spduHeader[idx] = 0x02
	idx++
	// flags: (.... .... .... ..1. = Duplex functional unit: True)
	spduHeader[idx] = 0x00
	idx++
	spduHeader[idx] = 0x02
	idx++

	// Calling Session Selector
	// Parameter type: Calling Session Selector (51)
	spduHeader[idx] = 0x33
	idx++
	// Parameter length
	spduHeader[idx] = 0x02
	idx++
	// Calling Session Selector
	spduHeader[idx] = sSelRemote[0]
	idx++
	spduHeader[idx] = sSelRemote[1]
	idx++

	// Called Session Selector
	// Parameter type: Called Session Selector (52)
	spduHeader[idx] = 0x34
	idx++
	// Parameter length
	spduHeader[idx] = 0x02
	idx++
	// Called Session Selector
	spduHeader[idx] = sSelLocal[0]
	idx++
	spduHeader[idx] = sSelLocal[1]
	idx++

	// Session user data
	// Parameter type: Session user data (193)
	spduHeader[idx] = 0xc1
	idx++
	// Parameter length
	spduHeader[idx] = (byte)(ssduLength & 0xff)
	// write session user data
	ssduList = append([][]byte{spduHeader}, ssduList...)

	ssduOffsets = append([]int{0}, ssduOffsets...)
	ssduLengths = append([]int{len(spduHeader)}, ssduLengths...)

	a.tConnection = tSAP.connectTo(address, port)

	a.tConnection.send(ssduList, ssduOffsets, ssduLengths)

	// TODO how much should be allocated here?
	pduBuffer := bytes.NewBuffer(make([]byte, 500))
	defer func() {
		r := recover()
		if r != nil {
			Throw("ResponseTimeout waiting for connection response.")
		}
	}()
	a.tConnection.receive(pduBuffer)

	// read ISO 8327-1 Header
	// SPDU Type: ACCEPT (14)
	spduType, err := pduBuffer.ReadByte()
	if err != nil {
		panic(err)
	}
	if spduType != 0x0e {
		Throw("ISO 8327-1 header wrong SPDU type, expected ACCEPT (14), got ", getSPDUTypeString(spduType), " (", string(spduType), ")")
	}
	_, _ = pduBuffer.ReadByte() // skip length byte

parameterLoop:
	for {
		// read parameter type
		b, err := pduBuffer.ReadByte()
		if err != nil {
			panic(err)
		}
		parameterType := b & 0xff
		// read parameter length
		b, err = pduBuffer.ReadByte()
		if err != nil {
			panic(err)
		}
		parameterLength := b & 0xff

		switch parameterType {
		// Connect Accept Item (5)
		case 0x05:

			bytesToRead := parameterLength
			for bytesToRead > 0 {
				// read parameter type
				caParameterType, err := pduBuffer.ReadByte()
				if err != nil {
					panic(err)
				}
				// read parameter length
				// int ca_parameterLength = res[idx++];
				_, err = pduBuffer.ReadByte()
				if err != nil {
					panic(err)
				}
				bytesToRead -= 2

				switch caParameterType & 0xff {
				// Protocol Options (19)
				case 0x13:
					// flags: .... ...0 = Able to receive extended
					// concatenated SPDU: False

					protocolOptions, err := pduBuffer.ReadByte()
					if err != nil {
						panic(err)
					}
					if protocolOptions != 0x00 {
						Throw("SPDU Connect Accept Item/Protocol Options is ", string(protocolOptions), ", expected 0")
					}
					bytesToRead--
					break
				// Version Number
				case 0x16:
					// flags .... ..1. = Protocol Version 2: True
					versionNumber, err := pduBuffer.ReadByte()
					if err != nil {
						panic(err)
					}
					if versionNumber != 0x02 {
						Throw("SPDU Connect Accept Item/Version Number is ", string(versionNumber), ", expected 2")
					}
					bytesToRead--
					break
				default:
					Throw("SPDU Connect Accept Item: parameter not implemented: ", string(caParameterType))
				}
			}
			break
		// Session Requirement (20)
		case 0x14:
			// flags: (.... .... .... ..1. = Duplex functional unit: True)

			sessionRequirement := a.extractInteger(pduBuffer, parameterLength)
			if sessionRequirement != 0x02 {
				Throw(
					"SPDU header parameter 'Session Requirement (20)' is ", strconv.FormatInt(sessionRequirement, 10), ", expected 2")
			}
			break
		// Calling Session Selector (51)
		case 0x33:
			css := a.extractInteger(pduBuffer, parameterLength)
			if css != 0x01 {
				Throw("SPDU header parameter 'Calling Session Selector (51)' is ", strconv.FormatInt(css, 10), ", expected 1")
			}
			break
		// Called Session Selector (52)
		case 0x34:

			calledSessionSelector := a.extractInteger(pduBuffer, parameterLength)
			if calledSessionSelector != 0x01 {
				Throw("SPDU header parameter 'Called Session Selector (52)' is ", strconv.FormatInt(calledSessionSelector, 10), ", expected 1")
			}
			break
		// Session user data (193)
		case 0xc1:
			break parameterLoop
		default:
			Throw("SPDU header parameter type ", string(parameterType), " not implemented")
		}
	}

	// got correct ACCEPT (AC) from the server

	a.connected = true

	return pduBuffer
}

func (a *AcseAssociation) extractInteger(buffer *bytes.Buffer, size byte) int64 {
	t := make([]byte, size)
	_, err := buffer.Read(t)
	if err != nil {
		panic(err)
	}
	switch size {
	case 1:
		return int64(t[0])
	case 2:
		return int64(binary.LittleEndian.Uint16(t))
	case 4:
		return int64(binary.LittleEndian.Uint32(t))

	case 8:
		return int64(binary.LittleEndian.Uint64(t))

	default:
		Throw("invalid length for reading numeric code")
	}
	return -1
}

func getSPDUTypeString(spduType byte) string {
	switch spduType {
	case 0:
		return "EXCEPTION REPORT (ER)"
	case 1:
		return "DATA TRANSFER (DT)"
	case 2:
		return "PLEASE TOKENS (PT)"
	case 5:
		return "EXPEDITED (EX)"
	case 7:
		return "PREPARE (PR)"
	case 8:
		return "NOT FINISHED (NF)"
	case 9:
		return "FINISH (FN)"
	case 10:
		return "DISCONNECT (DN)"
	case 12:
		return "REFUSE (RF)"
	case 13:
		return "CONNECT (CN)"
	case 14:
		return "ACCEPT (AC)"
	case 15:
		return "CONNECT DATA OVERFLOW (CDO)"
	case 16:
		return "OVERFLOW ACCEPT (OA)"
	case 21:
		return "GIVE TOKENS CONFIRM (GTC)"
	case 22:
		return "GIVE TOKENS ACK (GTA)"
	case 25:
		return "ABORT (AB)"
	case 26:
		return "ABORT ACCEPT (AA)"
	case 29:
		return "ACTIVITY RESUME (AR)"
	case 33:
		return "TYPED DATA (TD)"
	case 34:
		return "RESYNCHRONIZE ACK (RA)"
	case 41:
		return "MAJOR SYNC POINT (MAP)"
	case 42:
		return "MAJOR SYNC ACK (MAA)"
	case 45:
		return "ACTIVITY START (AS)"
	case 48:
		return "EXCEPTION DATA (ED)"
	case 49:
		return "MINOR SYNC POINT (MIP)"
	case 50:
		return "MINOR SYNC ACK (MIA)"
	case 53:
		return "RESYNCHRONIZE (RS)"
	case 57:
		return "ACTIVITY DISCARD (AD)"
	case 58:
		return "ACTIVITY DISCARD ACK (ADA)"
	case 61:
		return "CAPABILITY DATA (CD)"
	case 62:
		return "CAPABILITY DATA ACK (CDA)"
	case 64:
		return "UNIT DATA (UD)"
	default:
		return "<unknown SPDU type>"
	}
}

func getPresentationUserDataField(userDataBytes []byte) *UserData {
	presDataValues := NewPDVListPresentationDataValues()
	presDataValues.SingleASN1Type = NewBerAny(userDataBytes)
	pdvList := NewPDVList()
	pdvList.PresentationContextIdentifier = NewPresentationContextIdentifier()
	pdvList.PresentationDataValues = presDataValues

	fullyEncodedData := NewFullyEncodedData()
	pdvListList := fullyEncodedData.getPDVList()
	pdvListList = append(pdvListList, pdvList)

	userData := NewUserData()
	userData.FullyEncodedData = fullyEncodedData
	return userData
}

func NewAcseAssociation(tConnection *TConnection, pSelLocal []byte) *AcseAssociation {
	a := &AcseAssociation{}
	a.tConnection = tConnection
	a.pSelLocalBerOctetString = NewRespondingPresentationSelector(pSelLocal)

	return a
}
