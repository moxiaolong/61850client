package src

import (
	"bytes"
	"strconv"
)

type InformationReport struct {
	variableAccessSpecification *VariableAccessSpecification
	listOfAccessResult          *ListOfAccessResult
	tag                         *BerTag
	code                        []byte
}

func (r *InformationReport) decode(is *bytes.Buffer, withTag bool) int {

	numDecodedBytes := 0
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	if withTag {
		tlByteCount += r.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	r.variableAccessSpecification = NewVariableAccessSpecification()
	numDecodedBytes = r.variableAccessSpecification.decode(is, berTag)
	if numDecodedBytes != 0 {
		vByteCount += numDecodedBytes
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}
	if berTag.equals(128, 32, 0) {
		r.listOfAccessResult = NewListOfAccessResult()
		vByteCount += r.listOfAccessResult.decode(is, false)
		if lengthVal >= 0 && vByteCount == lengthVal {
			return tlByteCount + vByteCount
		}
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}

	if lengthVal < 0 {
		if !berTag.equals(0, 0, 0) {
			throw("Decoded sequence has wrong end of contents octets")
		}
		vByteCount += readEocByte(is)
		return tlByteCount + vByteCount
	}

	throw("Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func (r *InformationReport) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	codeLength += r.listOfAccessResult.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
	reverseOS.writeByte(0xA0)
	codeLength += 1

	codeLength += r.variableAccessSpecification.encode(reverseOS)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += r.tag.encode(reverseOS)
	}

	return codeLength
}

func NewInformationReport() *InformationReport {
	return &InformationReport{tag: NewBerTag(0, 32, 16)}
}
