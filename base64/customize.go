package base64

import (
	"encoding/base64"
)

const encoder = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

// base64 encode
func DiyEncode(str string) string {
	return base64.NewEncoding(encoder).EncodeToString([]byte(str))
}

// base64 decode
func StdDecode(str string) (string, error) {
	s, e := base64.NewEncoding(encoder).DecodeString(str)
	return string(s), e
}
