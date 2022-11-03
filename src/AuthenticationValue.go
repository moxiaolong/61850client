package src

import "bytes"

type AuthenticationValue struct {
}

func (v AuthenticationValue) decode(is *bytes.Buffer, t interface{}) int {

}

func (v AuthenticationValue) encode(os *ReverseByteArrayOutputStream) int {

}

func NewAuthenticationValue() *AuthenticationValue {
	return &AuthenticationValue{}
}
