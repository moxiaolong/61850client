package src

import "bytes"

type RejectReason struct {
	confirmedRequestPDU  *BerInteger
	confirmedResponsePDU *BerInteger
	confirmedErrorPDU    *BerInteger
	unconfirmedPDU       *BerInteger
	pduError             *BerInteger
	cancelRequestPDU     *BerInteger
	cancelResponsePDU    *BerInteger
	cancelErrorPDU       *BerInteger
	concludeRequestPDU   *BerInteger
	concludeResponsePDU  *BerInteger
	concludeErrorPDU     *BerInteger
	code                 []byte
}

func (r *RejectReason) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := (berTag != nil)

	if berTag == nil {
		berTag = NewEmptyBerTag()
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 0, 1) {
		r.confirmedRequestPDU = NewBerInteger(nil, 0)
		tlvByteCount += r.confirmedRequestPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 2) {
		r.confirmedResponsePDU = NewBerInteger(nil, 0)
		tlvByteCount += r.confirmedResponsePDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 3) {
		r.confirmedErrorPDU = NewBerInteger(nil, 0)
		tlvByteCount += r.confirmedErrorPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 4) {
		r.unconfirmedPDU = NewBerInteger(nil, 0)
		tlvByteCount += r.unconfirmedPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 5) {
		r.pduError = NewBerInteger(nil, 0)
		tlvByteCount += r.pduError.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 6) {
		r.cancelRequestPDU = NewBerInteger(nil, 0)
		tlvByteCount += r.cancelRequestPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 7) {
		r.cancelResponsePDU = NewBerInteger(nil, 0)
		tlvByteCount += r.cancelResponsePDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 8) {
		r.cancelErrorPDU = NewBerInteger(nil, 0)
		tlvByteCount += r.cancelErrorPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 9) {
		r.concludeRequestPDU = NewBerInteger(nil, 0)
		tlvByteCount += r.concludeRequestPDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 10) {
		r.concludeResponsePDU = NewBerInteger(nil, 0)
		tlvByteCount += r.concludeResponsePDU.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 11) {
		r.concludeErrorPDU = NewBerInteger(nil, 0)
		tlvByteCount += r.concludeErrorPDU.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (r *RejectReason) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if r.code != nil {
		reverseOS.write(r.code)
		return len(r.code)
	}

	codeLength := 0
	if r.concludeErrorPDU != nil {
		codeLength += r.concludeErrorPDU.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 11
		reverseOS.writeByte(0x8B)
		codeLength += 1
		return codeLength
	}

	if r.concludeResponsePDU != nil {
		codeLength += r.concludeResponsePDU.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 10
		reverseOS.writeByte(0x8A)
		codeLength += 1
		return codeLength
	}

	if r.concludeRequestPDU != nil {
		codeLength += r.concludeRequestPDU.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 9
		reverseOS.writeByte(0x89)
		codeLength += 1
		return codeLength
	}

	if r.cancelErrorPDU != nil {
		codeLength += r.cancelErrorPDU.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 8
		reverseOS.writeByte(0x88)
		codeLength += 1
		return codeLength
	}

	if r.cancelResponsePDU != nil {
		codeLength += r.cancelResponsePDU.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 7
		reverseOS.writeByte(0x87)
		codeLength += 1
		return codeLength
	}

	if r.cancelRequestPDU != nil {
		codeLength += r.cancelRequestPDU.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 6
		reverseOS.writeByte(0x86)
		codeLength += 1
		return codeLength
	}

	if r.pduError != nil {
		codeLength += r.pduError.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 5
		reverseOS.writeByte(0x85)
		codeLength += 1
		return codeLength
	}

	if r.unconfirmedPDU != nil {
		codeLength += r.unconfirmedPDU.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 4
		reverseOS.writeByte(0x84)
		codeLength += 1
		return codeLength
	}

	if r.confirmedErrorPDU != nil {
		codeLength += r.confirmedErrorPDU.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 3
		reverseOS.writeByte(0x83)
		codeLength += 1
		return codeLength
	}

	if r.confirmedResponsePDU != nil {
		codeLength += r.confirmedResponsePDU.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 2
		reverseOS.writeByte(0x82)
		codeLength += 1
		return codeLength
	}

	if r.confirmedRequestPDU != nil {
		codeLength += r.confirmedRequestPDU.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 1
		reverseOS.writeByte(0x81)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewRejectReason() *RejectReason {
	return &RejectReason{}
}
