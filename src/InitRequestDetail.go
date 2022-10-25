package src

type InitRequestDetail struct {
	servicesSupportedCalling *ServiceSupportOptions
	proposedParameterCBB     *ParameterSupportOptions
	proposedVersionNumber    *Integer16
	Tag                      *BerTag
}

func (d *InitRequestDetail) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	codeLength := 0
	codeLength += d.servicesSupportedCalling.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 2
	reverseOS.writeByte(0x82)
	codeLength += 1

	codeLength += d.proposedParameterCBB.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 1
	reverseOS.writeByte(0x81)
	codeLength += 1

	codeLength += d.proposedVersionNumber.encode(reverseOS, false)
	// write tag: CONTEXT_CLASS, PRIMITIVE, 0
	reverseOS.writeByte(0x80)
	codeLength += 1

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += d.Tag.encode(reverseOS)
	}

	return codeLength

}

func NewInitRequestDetail() *InitRequestDetail {
	return &InitRequestDetail{Tag: NewBerTag(0, 32, 16)}
}
