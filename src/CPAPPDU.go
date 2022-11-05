package src

import (
	"bytes"
	"strconv"
)

type CPAPPDU struct {
	tag                  *BerTag
	modeSelector         *ModeSelector
	normalModeParameters *NormalModeParameters
}

func (c *CPAPPDU) encode()  {

	if (code != nil) {
		reverseOS.write(code);
		if (withTag) {
			return tag.encode(reverseOS) + code.length;
		}
		return code.length;
	}

	int codeLength = 0;
	if (normalModeParameters != nil) {
		codeLength += normalModeParameters.encode(reverseOS, false);
		// write tag: CONTEXT_CLASS, CONSTRUCTED, 2
		reverseOS.write(0xA2);
		codeLength += 1;
	}

	codeLength += modeSelector.encode(reverseOS, false);
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
	reverseOS.write(0xA0);
	codeLength += 1;

	codeLength += BerLength.encodeLength(reverseOS, codeLength);

	if (withTag) {
		codeLength += tag.encode(reverseOS);
	}

	return codeLength;
}

func (c *CPAPPDU) decode(is *bytes.Buffer) int {
	tlByteCount := 0
	vByteCount := 0
	berTag := NewBerTag(0, 0, 0)

	tlByteCount += c.tag.decodeAndCheck(is)

	length := NewBerLength()
	tlByteCount += length.decode(is)

	lengthVal := length.val

	for vByteCount < lengthVal || lengthVal < 0 {
		vByteCount += berTag.decode(is)
		if berTag.equals(128, 32, 0) {
			c.modeSelector = NewModeSelector()
			vByteCount += c.modeSelector.decode(is, false)
		} else if berTag.equals(128, 32, 2) {
			c.normalModeParameters = NewNormalModeParameters()
			vByteCount += c.normalModeParameters.decode(is, false)
		} else if lengthVal < 0 && berTag.equals(0, 0, 0) {
			vByteCount += readEocByte(is)
			return tlByteCount + vByteCount
		} else {
			throw("tag does not match any set component: ", berTag.toString())
		}
	}
	if vByteCount != lengthVal {
		throw("Length of set does not match length tag, length tag: ", strconv.Itoa(lengthVal), ", actual set length: ", strconv.Itoa(vByteCount))
	}
	return tlByteCount + vByteCount
}

func NewCPAPPDU() *CPAPPDU {
	return &CPAPPDU{tag: NewBerTag(0, 32, 17)}
}
