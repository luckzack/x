package sha256

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func Encrypt(b []byte) string {
	r := sha256.Sum256(b)
	return hex.EncodeToString(r[:])
}

func EncryptString(s string) string {
	r := sha256.Sum256([]byte(s))
	return hex.EncodeToString(r[:])
}

func EncryptFile(path string) string {
	f, e := os.Open(path)
	if e != nil {
		return ""
	}
	h := sha256.New()
	_, e = io.Copy(h, f)
	if e != nil {
		return ""
	}
	return hex.EncodeToString(h.Sum(nil))
}
