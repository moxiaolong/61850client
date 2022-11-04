package src

import (
	"bytes"
)

type AssociateSourceDiagnostic struct {
	acseServiceUser     *BerInteger
	acseServiceProvider *BerInteger
	code                []byte
}

func (d *AssociateSourceDiagnostic) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {
		length := NewBerLength()
		tlvByteCount += length.decode(is)
		d.acseServiceUser = NewBerInteger(nil, 0)
		tlvByteCount += d.acseServiceUser.decode(is, true)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 2) {
		length := NewBerLength()
		tlvByteCount += length.decode(is)
		d.acseServiceProvider = NewBerInteger(nil, 0)
		tlvByteCount += d.acseServiceProvider.decode(is, true)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding CHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (d *AssociateSourceDiagnostic) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if d.code != nil {
		reverseOS.write(d.code)
		return len(d.code)
	}

	codeLength := 0

	sublength := 0

	if d.acseServiceProvider != nil {
		sublength = d.acseServiceProvider.encode(reverseOS, true)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.writeByte(0xA2)
		codeLength += 1
		return codeLength
	}

	if d.acseServiceUser != nil {
		sublength = d.acseServiceUser.encode(reverseOS, true)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0xA1)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding CHOICE: No element of CHOICE was selected.")
	return 0
}

func NewAssociateSourceDiagnostic() *AssociateSourceDiagnostic {
	return &AssociateSourceDiagnostic{}
}
