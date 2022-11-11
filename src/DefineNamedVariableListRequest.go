package src

import (
	"bytes"
	"strconv"
)

type DefineNamedVariableListRequest struct {
	tag              *BerTag
	code             []byte
	listOfVariable   *VariableDefs
	variableListName *ObjectName
}

func (r *DefineNamedVariableListRequest) decode(is *bytes.Buffer, withTag bool) int {
	tlByteCount := 0
	vByteCount := 0
	numDecodedBytes := 0
	berTag := NewEmptyBerTag()

	if withTag {
		tlByteCount += r.tag.decodeAndCheck(is)
	}

	length := NewBerLength()
	tlByteCount += length.decode(is)
	lengthVal := length.val
	vByteCount += berTag.decode(is)

	r.variableListName = NewObjectName()
	numDecodedBytes = r.variableListName.decode(is, berTag)
	if numDecodedBytes != 0 {
		vByteCount += numDecodedBytes
		vByteCount += berTag.decode(is)
	} else {
		throw("tag does not match mandatory sequence component.")
	}
	if berTag.equals(128, 32, 0) {
		r.listOfVariable = NewVariableDefs()
		vByteCount += r.listOfVariable.decode(is, false)
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

	throw(
		"Unexpected end of sequence, length tag: " + strconv.Itoa(lengthVal) + ", bytes decoded: " + strconv.Itoa(vByteCount))
	return 0
}

func (r *DefineNamedVariableListRequest) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if r.code != nil {
		reverseOS.write(r.code)
		if withTag {
			return r.tag.encode(reverseOS) + len(r.code)
		}
		return len(r.code)
	}

	codeLength := 0
	codeLength += r.listOfVariable.encode(reverseOS, false)
	// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 0
	reverseOS.writeByte(0xA0)
	codeLength += 1

	codeLength += r.variableListName.encode(reverseOS)

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += r.tag.encode(reverseOS)
	}

	return codeLength
}

func NewDefineNamedVariableListRequest() *DefineNamedVariableListRequest {
	return &DefineNamedVariableListRequest{tag: NewBerTag(0, 32, 16)}
}
