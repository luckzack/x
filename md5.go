package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"crypto/sha1"
)

func Md5(raw string) string {
	h := md5.New()
	io.WriteString(h, raw)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha1(raw string) string {
	h := sha1.New()
	io.WriteString(h, raw)

	return fmt.Sprintf("%x", h.Sum(nil))
}
