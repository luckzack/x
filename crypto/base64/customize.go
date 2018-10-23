package base64

import (
	"encoding/base64"
)

const my_encoder = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789[>"

// base64 encode
func DiyEncode(str string, encoders ...string) string {
	if len(encoders) > 0 {
		return base64.NewEncoding(encoders[0]).EncodeToString([]byte(str))
	} else {
		return base64.NewEncoding(my_encoder).EncodeToString([]byte(str))
	}

}

func DiyEncodeBytes(v []byte, encoders ...string) string {
	if len(encoders) > 0 {
		return base64.NewEncoding(encoders[0]).EncodeToString(v)
	} else {
		return base64.NewEncoding(my_encoder).EncodeToString(v)
	}
}

// base64 decode
func DiyDecode(str string, encoders ...string) (string, error) {

	if len(encoders) > 0 {
		s, e := base64.NewEncoding(encoders[0]).DecodeString(str)
		return string(s), e
	} else {
		s, e := base64.NewEncoding(my_encoder).DecodeString(str)
		return string(s), e
	}

}

func DiyDecodeToBytes(str string, encoders ...string) ([]byte, error) {
	if len(encoders) > 0 {
		return base64.NewEncoding(encoders[0]).DecodeString(str)
	} else {
		return base64.NewEncoding(my_encoder).DecodeString(str)
	}

}
