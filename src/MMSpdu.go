package src

type MMSpdu struct {
	InitiateErrorPDU    *InitiateErrorPDU
	InitiateResponsePDU *InitiateResponsePDU
}

func newMMSpdu() *MMSpdu {
	return nil
}
func (s *MMSpdu) encode(stream *ReverseByteArrayOutputStream) {

}

func (s *MMSpdu) decode(stream *ByteBufferInputStream) {

}

func constructInitRequestPdu(int, int, int, int, []byte) *MMSpdu {
	return &MMSpdu{}
}
