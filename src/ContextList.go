package src

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
