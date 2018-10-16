package base64

import (
	"encoding/base64"
)

// base64 encode
func StdEncode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func StdEncodeBytes(v []byte) string {
	return base64.StdEncoding.EncodeToString(v)
}

// base64 decode
func StdDecode(str string) (string, error) {
	s, e := base64.StdEncoding.DecodeString(str)
	return string(s), e
}

func StdDecodeToBytes(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
