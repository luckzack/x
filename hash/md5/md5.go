package md5

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func Encrypt(b []byte) string {
	h := md5.New()
	h.Write(b)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func EncryptString(v string) string {
	h := md5.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func EncryptFile(path string) string {
	f, e := os.Open(path)
	if e != nil {
		return ""
	}
	h := md5.New()
	_, e = io.Copy(h, f)
	if e != nil {
		return ""
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
