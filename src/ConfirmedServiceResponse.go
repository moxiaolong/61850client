package src

import "bytes"

type ConfirmedServiceResponse struct {
	code                           []byte
	getVariableAccessAttributes    *GetVariableAccessAttributesResponse
	getNameList                    *GetNameListResponse
	read                           *ReadResponse
	writeByteByte                  *WriteResponse
	defineNamedVariableList        *DefineNamedVariableListResponse
	getNamedVariableListAttributes *GetNamedVariableListAttributesResponse
	deleteNamedVariableList        *DeleteNamedVariableListResponse
	fileOpen                       *FileOpenResponse
	fileRead                       *FileReadResponse
	fileClose                      *FileCloseResponse
	fileDelete                     *FileDeleteResponse
	fileDirectory                  *FileDirectoryResponse
}

func (r *ConfirmedServiceResponse) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewBerTag(0, 0, 0)
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {
		r.getNameList = NewGetNameListResponse()
		tlvByteCount += r.getNameList.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 4) {
		r.read = NewReadResponse()
		tlvByteCount += r.read.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 5) {
		r.writeByteByte = NewWriteResponse()
		tlvByteCount += r.writeByteByte.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 6) {
		r.getVariableAccessAttributes = NewGetVariableAccessAttributesResponse()
		tlvByteCount += r.getVariableAccessAttributes.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 11) {
		r.defineNamedVariableList = NewDefineNamedVariableListResponse()
		tlvByteCount += r.defineNamedVariableList.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 12) {
		r.getNamedVariableListAttributes = NewGetNamedVariableListAttributesResponse()
		tlvByteCount += r.getNamedVariableListAttributes.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 13) {
		r.deleteNamedVariableList = NewDeleteNamedVariableListResponse()
		tlvByteCount += r.deleteNamedVariableList.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 72) {
		r.fileOpen = NewFileOpenResponse()
		tlvByteCount += r.fileOpen.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 73) {
		r.fileRead = NewFileReadResponse()
		tlvByteCount += r.fileRead.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 74) {
		r.fileClose = NewFileCloseResponse()
		tlvByteCount += r.fileClose.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 76) {
		r.fileDelete = NewFileDeleteResponse()
		tlvByteCount += r.fileDelete.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 77) {
		r.fileDirectory = NewFileDirectoryResponse()
		tlvByteCount += r.fileDirectory.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (r *ConfirmedServiceResponse) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if r.code != nil {
		reverseOS.write(r.code)
		return len(r.code)
	}

	codeLength := 0
	if r.fileDirectory != nil {
		codeLength += r.fileDirectory.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 77
		reverseOS.writeByte(0x4D)
		reverseOS.writeByte(0xBF)
		codeLength += 2
		return codeLength
	}

	if r.fileDelete != nil {
		codeLength += r.fileDelete.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 76
		reverseOS.writeByte(0x4C)
		reverseOS.writeByte(0x9F)
		codeLength += 2
		return codeLength
	}

	if r.fileClose != nil {
		codeLength += r.fileClose.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 74
		reverseOS.writeByte(0x4A)
		reverseOS.writeByte(0x9F)
		codeLength += 2
		return codeLength
	}

	if r.fileRead != nil {
		codeLength += r.fileRead.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 73
		reverseOS.writeByte(0x49)
		reverseOS.writeByte(0xBF)
		codeLength += 2
		return codeLength
	}

	if r.fileOpen != nil {
		codeLength += r.fileOpen.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 72
		reverseOS.writeByte(0x48)
		reverseOS.writeByte(0xBF)
		codeLength += 2
		return codeLength
	}

	if r.deleteNamedVariableList != nil {
		codeLength += r.deleteNamedVariableList.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 13
		reverseOS.writeByte(0xAD)
		codeLength += 1
		return codeLength
	}

	if r.getNamedVariableListAttributes != nil {
		codeLength += r.getNamedVariableListAttributes.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 12
		reverseOS.writeByte(0xAC)
		codeLength += 1
		return codeLength
	}

	if r.defineNamedVariableList != nil {
		codeLength += r.defineNamedVariableList.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 11
		reverseOS.writeByte(0x8B)
		codeLength += 1
		return codeLength
	}

	if r.getVariableAccessAttributes != nil {
		codeLength += r.getVariableAccessAttributes.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 6
		reverseOS.writeByte(0xA6)
		codeLength += 1
		return codeLength
	}

	if r.writeByteByte != nil {
		codeLength += r.writeByteByte.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 5
		reverseOS.writeByte(0xA5)
		codeLength += 1
		return codeLength
	}

	if r.read != nil {
		codeLength += r.read.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 4
		reverseOS.writeByte(0xA4)
		codeLength += 1
		return codeLength
	}

	if r.getNameList != nil {
		codeLength += r.getNameList.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 1
		reverseOS.writeByte(0xA1)
		codeLength += 1
		return codeLength
	}

	throw("Error encoding WriteResponseCHOICE: No element of WriteResponseCHOICE was selected.")
	return 0
}

func NewConfirmedServiceResponse() *ConfirmedServiceResponse {
	return &ConfirmedServiceResponse{}
}
