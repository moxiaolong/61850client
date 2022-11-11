package src

type ProtocolVersion struct {
	BerBitString
}

func NewProtocolVersion() *ProtocolVersion {
	return &ProtocolVersion{BerBitString: *NewBerBitString(nil, nil, 0)}
}
