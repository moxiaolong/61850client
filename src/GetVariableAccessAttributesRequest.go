package src

import "bytes"

type GetVariableAccessAttributesRequest struct {
	name *ObjectName
	code []byte
}

func (r *GetVariableAccessAttributesRequest) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := (berTag != nil)

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 0) {
		length := NewBerLength()
		tlvByteCount += length.decode(is)
		r.name = NewObjectName()
		tlvByteCount += r.name.decode(is, nil)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: Tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (r *GetVariableAccessAttributesRequest) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if r.code != nil {
		reverseOS.write(r.code)
		return len(r.code)
	}

	codeLength := 0
	sublength := 0

	if r.name != nil {
		sublength = r.name.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
		reverseOS.writeByte(0xA0)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewGetVariableAccessAttributesRequest() *GetVariableAccessAttributesRequest {
	return &GetVariableAccessAttributesRequest{}
}
