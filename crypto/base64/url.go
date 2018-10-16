package base64

import (
	"encoding/base64"
)

// base64 encode
func UrlEncode(str string) string {
	return base64.URLEncoding.EncodeToString([]byte(str))
}

func UrlEncodeBytes(v []byte) string {
	return base64.URLEncoding.EncodeToString(v)
}

// base64 decode
func UrlDecode(str string) (string, error) {
	s, e := base64.URLEncoding.DecodeString(str)
	return string(s), e
}

func UrlDecodeToBytes(str string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(str)
}
