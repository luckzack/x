package base64

import (
	"encoding/base64"
)

const encoder = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789[!"

// base64 encode
func DiyEncode(str string) string {
	return base64.NewEncoding(encoder).EncodeToString([]byte(str))
}

func DiyEncodeBytes(v []byte) string {
	return base64.NewEncoding(encoder).EncodeToString(v)
}

// base64 decode
func DiyDecode(str string) (string, error) {
	s, e := base64.NewEncoding(encoder).DecodeString(str)
	return string(s), e
}

func DiyDecodeToBytes(str string) ([]byte, error) {
	return base64.NewEncoding(encoder).DecodeString(str)
}
