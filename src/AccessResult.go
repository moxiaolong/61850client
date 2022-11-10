package src

import "bytes"

type AccessResult struct {
	success *Data
	failure *DataAccessError
	code    []byte
}

func (r *AccessResult) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := (berTag != nil)

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	numDecodedBytes := 0

	if berTag.equals(128, 0, 0) {
		r.failure = NewDataAccessError()
		tlvByteCount += r.failure.decode(is, false)
		return tlvByteCount
	}

	r.success = NewData()
	numDecodedBytes = r.success.decode(is, berTag)
	if numDecodedBytes != 0 {
		return tlvByteCount + numDecodedBytes
	} else {
		r.success = nil
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (r *AccessResult) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if r.code != nil {
		reverseOS.write(r.code)
		return len(r.code)
	}

	codeLength := 0
	if r.success != nil {
		codeLength += r.success.encode(reverseOS)
		return codeLength
	}

	if r.failure != nil {
		codeLength += r.failure.encode(reverseOS, false)
		// write tag: CONTEXT_CLASS, PRIMITIVE, 0
		reverseOS.writeByte(0x80)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewAccessResult() *AccessResult {
	return &AccessResult{}
}
