package src

type CPType struct {
	ModeSelector         *ModeSelector
	NormalModeParameters *CPTypeNormalModeParameters
}

func (t CPType) encode(stream *ReverseByteArrayOutputStream, b bool) {

}

func NewCPType() *CPType {
	return &CPType{}
}
