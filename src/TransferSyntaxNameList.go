package src

type TransferSyntaxNameList struct {
	tag   *BerTag
	seqOf []*TransferSyntaxName
}

func (l *TransferSyntaxNameList) encode(reverseOS *ReverseByteArrayOutputStream, withTag bool) int {
	if code != nil {
		reverseOS.write(code)
		if withTag {
			return tag.encode(reverseOS) + code.length
		}
		return code.length
	}

	codeLength := 0
	for i := len(l.seqOf) - 1; i >= 0; i-- {
		codeLength += l.seqOf[i].encode(reverseOS, true)
	}

	codeLength += encodeLength(reverseOS, codeLength)

	if withTag {
		codeLength += l.tag.encode(reverseOS)
	}

	return codeLength
}
func (l *TransferSyntaxNameList) decode() {
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

		if (!berTag.equals(TransferSyntaxName.tag)) {
			throw("Tag does not match mandatory sequence of/set of component.");
		}
		TransferSyntaxName element = NewTransferSyntaxName();
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

func NewTransferSyntaxNameList() *TransferSyntaxNameList {
	return &TransferSyntaxNameList{tag: NewBerTag(0, 32, 16)}
}
