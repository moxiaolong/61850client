package src

import "bytes"

type AcseAssociation struct {
	MessageTimeout int
}

func (a *AcseAssociation) getAssociateResponseAPdu() *bytes.Buffer {
	return nil
}

func (a *AcseAssociation) disconnect() {

}
