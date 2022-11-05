package src

import "bytes"

type ContextList struct {
	code  []byte
	tag   *BerTag
	seqOf []*SEQUENCE
}

func NewContextList(code []byte) *ContextList {
	return &ContextList{code: code, tag: NewBerTag(0, 32, 16)}
}

func (c *ContextList) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if c.code != nil {
		reverseOS.write(c.code)
		if withTag {
			return c.tag.encode(reverseOS) + len(c.code)
		}
		return len(c.code)
	}

	codeLength := 0
	for i := len(c.seqOf) - 1; i >= 0; i-- {
		codeLength += c.seqOf[i].encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += c.tag.encode(reverseOS)
	}

	return codeLength
}

func (c *ContextList) decode(is *bytes.Buffer, withTag bool) int {
	int tlByteCount = 0;
	int vByteCount = 0;
	BerTag berTag = NewBerTag(0,0,0);
	if (withTag) {
		tlByteCount += tag.decodeAndCheck(is);
	}

	BerLength length = NewBerLength();
	tlByteCount += length.decode(is);
	int lengthVal = length.val;

	while (vByteCount < lengthVal || lengthVal < 0) {
		vByteCount += berTag.decode(is);

		if (lengthVal < 0 && berTag.equals(0, 0, 0)) {
			vByteCount += BerLength.readEocByte(is);
			break;
		}

		if (!berTag.equals(SEQUENCE.tag)) {
			throw("Tag does not match mandatory sequence of/set of component.");
		}
		SEQUENCE element = NewSEQUENCE();
		vByteCount += element.decode(is, false);
		seqOf.add(element);
	}
	if (lengthVal >= 0 && vByteCount != lengthVal) {
		throw(
			"Decoded SequenceOf or SetOf has wrong length. Expected "
		+ lengthVal
		+ " but has "
		+ vByteCount);
	}
	return tlByteCount + vByteCount;
}
