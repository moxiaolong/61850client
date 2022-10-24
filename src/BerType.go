package src

type BerType interface {
	encode(*ReverseByteArrayOutputStream) int
}
