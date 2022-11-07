package src

import "bytes"

type RejectReason struct {
}

func (r RejectReason) decode(is *bytes.Buffer, tag *BerTag) int {

}

func (r RejectReason) encode(os *ReverseByteArrayOutputStream) int {

}

func NewRejectReason() *RejectReason {
	return &RejectReason{}
}
