package src

type ProtocolVersion struct {
	BerBitString
}

func NewProtocolVersion() *ProtocolVersion {
	return &ProtocolVersion{}
}
