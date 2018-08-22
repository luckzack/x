package sha1

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
)

func Encrypt(b []byte) string {
	r := sha1.Sum(b)
	return hex.EncodeToString(r[:])
}

func EncryptString(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func EncryptFile(path string) string {
	f, e := os.Open(path)
	if e != nil {
		return ""
	}
	h := sha1.New()
	_, e = io.Copy(h, f)
	if e != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
