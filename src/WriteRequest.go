package src

import "bytes"

type WriteRequest struct {
}

func (r *WriteRequest) decode(is *bytes.Buffer, b bool) int {
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
	vByteCount += berTag.decode(is);

	variableAccessSpecification = new VariableAccessSpecification();
	numDecodedBytes = variableAccessSpecification.decode(is, berTag);
	if (numDecodedBytes != 0) {
		vByteCount += numDecodedBytes;
		vByteCount += berTag.decode(is);
	} else {
		throw new IOException("Tag does not match mandatory sequence component.");
	}
	if (berTag.equals(BerTag.CONTEXT_CLASS, BerTag.CONSTRUCTED, 0)) {
		listOfData = new ListOfData();
		vByteCount += listOfData.decode(is, false);
		if (lengthVal >= 0 && vByteCount == lengthVal) {
			return tlByteCount + vByteCount;
		}
		vByteCount += berTag.decode(is);
	} else {
		throw new IOException("Tag does not match mandatory sequence component.");
	}

	if (lengthVal < 0) {
		if (!berTag.equals(0, 0, 0)) {
			throw new IOException("Decoded sequence has wrong end of contents octets");
		}
		vByteCount += BerLength.readEocByte(is);
		return tlByteCount + vByteCount;
	}

	throw new IOException(
		"Unexpected end of sequence, length tag: " + lengthVal + ", bytes decoded: " + vByteCount);
}

func (r *WriteRequest) encode(os *ReverseByteArrayOutputStream, b bool) int {
	if (code != null) {
		reverseOS.write(code);
		if (withTag) {
			return tag.encode(reverseOS) + code.length;
		}
		return code.length;
	}

	int codeLength = 0;
	codeLength += listOfData.encode(reverseOS, false);
	// write tag: CONTEXT_CLASS, CONSTRUCTED, 0
	reverseOS.write(0xA0);
	codeLength += 1;

	codeLength += variableAccessSpecification.encode(reverseOS);

	codeLength += BerLength.encodeLength(reverseOS, codeLength);

	if (withTag) {
		codeLength += tag.encode(reverseOS);
	}

	return codeLength;
}

func NewWriteRequest() *WriteRequest {
	return &WriteRequest{}
}
