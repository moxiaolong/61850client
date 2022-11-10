package src

import "bytes"

type WriteResponse struct {
}

func (r *WriteResponse) decode(is *bytes.Buffer, b bool) int {
	int tlByteCount = 0;
	int vByteCount = 0;
	int numDecodedBytes;
	BerTag berTag = new BerTag();
	if (withTag) {
		tlByteCount += tag.decodeAndCheck(is);
	}

	BerLength length = new BerLength();
	tlByteCount += length.decode(is);
	int lengthVal = length.val;

	while (vByteCount < lengthVal || lengthVal < 0) {
		vByteCount += berTag.decode(is);

		if (lengthVal < 0 && berTag.equals(0, 0, 0)) {
			vByteCount += BerLength.readEocByte(is);
			break;
		}

		CHOICE element = new CHOICE();
		numDecodedBytes = element.decode(is, berTag);
		if (numDecodedBytes == 0) {
			throw new IOException("Tag did not match");
		}
		vByteCount += numDecodedBytes;
		seqOf.add(element);
	}
	if (lengthVal >= 0 && vByteCount != lengthVal) {
		throw new IOException(
			"Decoded SequenceOf or SetOf has wrong length. Expected "
		+ lengthVal
		+ " but has "
		+ vByteCount);
	}
	return tlByteCount + vByteCount;
}

func (r *WriteResponse) encode(os *ReverseByteArrayOutputStream, b bool) int {
	if (code != null) {
		reverseOS.write(code);
		if (withTag) {
			return tag.encode(reverseOS) + code.length;
		}
		return code.length;
	}

	int codeLength = 0;
	for (int i = (seqOf.size() - 1); i >= 0; i--) {
		codeLength += seqOf.get(i).encode(reverseOS);
	}

	codeLength += BerLength.encodeLength(reverseOS, codeLength);

	if (withTag) {
		codeLength += tag.encode(reverseOS);
	}

	return codeLength;
}

func NewWriteResponse() *WriteResponse {
	return &WriteResponse{}
}
