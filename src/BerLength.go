package src

import (
	"bytes"
	"strconv"
)

type BerLength struct {
	val int
}

func readEocByte(is *bytes.Buffer) int {
	b, err := is.ReadByte()
	if err != nil {
		throw("Unexpected end of input stream.")
	}
	if b != 0 {
		throw("Byte ", string(b), " does not match end of contents octet of zero.")
	} else {
		return 1
	}
	return -1

}

func (l *BerLength) decode(is *bytes.Buffer) int {
	b, err := is.ReadByte()
	if err != nil {
		throw("Unexpected end of input stream.")
	}
	l.val = int(b)
	if l.val < 128 {

		return 1
	} else {

		lengthLength := l.val & 127
		if lengthLength == 0 {
			l.val = -1
			return 1
		} else if lengthLength > 4 {
			throw("Length is out of bounds: ", strconv.Itoa(lengthLength))
		} else {
			l.val = 0

			for i := 0; i < lengthLength; i++ {

				nextByte, err := is.ReadByte()
				if err != nil {
					throw("Unexpected end of input stream.")
				}

				l.val |= int(nextByte) << 8 * (lengthLength - i - 1)
			}

			return lengthLength + 1
		}
	}

	return 0
}

func (l *BerLength) readEocIfIndefinite(is *bytes.Buffer) int {
	if l.val >= 0 {
		return 0
	} else {
		readEocByte(is)
		readEocByte(is)
		return 2
	}
}

func NewBerLength() *BerLength {
	return &BerLength{}
}
