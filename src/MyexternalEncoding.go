package src

type MyexternalEncoding struct {
	SingleASN1Type *BerAny
}

func NewMyexternalEncoding() *MyexternalEncoding {
	return &MyexternalEncoding{}
}
