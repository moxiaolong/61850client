package src

import "bytes"

type InitResponseDetail struct {
	NegotiatedVersionNumber *NegotiatedVersionNumber
	ServicesSupportedCalled *ServicesSupportedCalled
	tag                     *BerTag
}

func (d InitResponseDetail) decode(is *bytes.Buffer, b bool) int {

}

func (d InitResponseDetail) encode(os *ReverseByteArrayOutputStream, b bool) int {

}

func NewInitResponseDetail() *InitResponseDetail {
	return &InitResponseDetail{tag: NewBerTag(0, 32, 16)}
}
