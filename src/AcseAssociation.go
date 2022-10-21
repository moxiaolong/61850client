package src

import (
	"bytes"
)

type AcseAssociation struct {
	MessageTimeout          int
	associateResponseAPDU   *bytes.Buffer
	tConnection             *TConnection
	pSelLocalBerOctetString *RespondingPresentationSelector
	tSap                    *ClientTSap
	connected               bool
}

func NewAcseAssociation(i []byte) *AcseAssociation {
	return &AcseAssociation{}
}

func (a *AcseAssociation) getAssociateResponseAPdu() *bytes.Buffer {
	returnBuffer := a.associateResponseAPDU
	a.associateResponseAPDU = nil
	return returnBuffer
}

func (a *AcseAssociation) disconnect() {

}

func (a *AcseAssociation) startAssociation(apdu *bytes.Buffer, address string, port int, i []byte, i2 []byte, i3 []byte, sap *ClientTSap, ints []int, ints2 []int, i4 int, i5 int) {

}
