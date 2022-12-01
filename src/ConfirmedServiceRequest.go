package src

import "bytes"

type ConfirmedServiceRequest struct {
	getNameList                    *GetNameListRequest
	read                           *ReadRequest
	write                          *WriteRequest
	getVariableAccessAttributes    *GetVariableAccessAttributesRequest
	defineNamedVariableList        *DefineNamedVariableListRequest
	getNamedVariableListAttributes *GetNamedVariableListAttributesRequest
	deleteNamedVariableList        *DeleteNamedVariableListRequest
	fileOpen                       *FileOpenRequest
	fileRead                       *FileReadRequest
	fileClose                      *FileCloseRequest
	fileDelete                     *FileDeleteRequest
	fileDirectory                  *FileDirectoryRequest
	code                           []byte
}

func (r *ConfirmedServiceRequest) decode(is *bytes.Buffer, berTag *BerTag) int {
	tlvByteCount := 0
	tagWasPassed := berTag != nil

	if berTag == nil {
		berTag = NewEmptyBerTag()
		tlvByteCount += berTag.decode(is)
	}

	if berTag.equals(128, 32, 1) {
		r.getNameList = NewGetNameListRequest()
		tlvByteCount += r.getNameList.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 4) {
		r.read = NewReadRequest()
		tlvByteCount += r.read.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 5) {
		r.write = NewWriteRequest()
		tlvByteCount += r.write.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 6) {
		length := NewBerLength()
		tlvByteCount += length.decode(is)
		r.getVariableAccessAttributes = NewGetVariableAccessAttributesRequest()
		tlvByteCount += r.getVariableAccessAttributes.decode(is, nil)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 11) {
		r.defineNamedVariableList = NewDefineNamedVariableListRequest()
		tlvByteCount += r.defineNamedVariableList.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 12) {
		length := NewBerLength()
		tlvByteCount += length.decode(is)
		r.getNamedVariableListAttributes = NewGetNamedVariableListAttributesRequest()
		tlvByteCount += r.getNamedVariableListAttributes.decode(is, nil)
		tlvByteCount += length.readEocIfIndefinite(is)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 13) {
		r.deleteNamedVariableList = NewDeleteNamedVariableListRequest()
		tlvByteCount += r.deleteNamedVariableList.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 72) {
		r.fileOpen = NewFileOpenRequest()
		tlvByteCount += r.fileOpen.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 73) {
		r.fileRead = NewFileReadRequest()
		tlvByteCount += r.fileRead.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 0, 74) {
		r.fileClose = NewFileCloseRequest()
		tlvByteCount += r.fileClose.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 76) {
		r.fileDelete = NewFileDeleteRequest()
		tlvByteCount += r.fileDelete.decode(is, false)
		return tlvByteCount
	}

	if berTag.equals(128, 32, 77) {
		r.fileDirectory = NewFileDirectoryRequest()
		tlvByteCount += r.fileDirectory.decode(is, false)
		return tlvByteCount
	}

	if tagWasPassed {
		return 0
	}

	throw("Error decoding WriteResponseCHOICE: tag " + berTag.toString() + " matched to no item.")
	return 0
}

func (r *ConfirmedServiceRequest) encode(reverseOS *ReverseByteArrayOutputStream) int {
	if r.code != nil {
		reverseOS.write(r.code)
		return len(r.code)
	}
	codeLength := 0
	sublength := 0

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
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 76
		reverseOS.writeByte(0x4C)
		reverseOS.writeByte(0xBF)
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
		// writeByte tag: CONTEXT_CLASS, PRIMITIVE, 73
		reverseOS.writeByte(0x49)
		reverseOS.writeByte(0x9F)
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
		sublength = r.getNamedVariableListAttributes.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 12
		reverseOS.writeByte(0xAC)
		codeLength += 1
		return codeLength
	}

	if r.defineNamedVariableList != nil {
		codeLength += r.defineNamedVariableList.encode(reverseOS, false)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 11
		reverseOS.writeByte(0xAB)
		codeLength += 1
		return codeLength
	}

	if r.getVariableAccessAttributes != nil {
		sublength = r.getVariableAccessAttributes.encode(reverseOS)
		codeLength += sublength
		codeLength += encodeLength(reverseOS, sublength)
		// writeByte tag: CONTEXT_CLASS, CONSTRUCTED, 6
		reverseOS.writeByte(0xA6)
		codeLength += 1
		return codeLength
	}

	if r.write != nil {
		codeLength += r.write.encode(reverseOS, false)
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

func NewConfirmedServiceRequest() *ConfirmedServiceRequest {
	return &ConfirmedServiceRequest{}
}
