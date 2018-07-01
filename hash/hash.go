package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
)

func Md5String(raw string) string {
	h := md5.New()
	io.WriteString(h, raw)

	return fmt.Sprintf("%x", h.Sum(nil))
}

func Md5sumInterface(raw interface{}) string {
	bytes, _ := json.Marshal(raw)

	h := md5.New()
	h.Write(bytes)
	defer func() {
		h = nil
	}()
	return hex.EncodeToString(h.Sum(nil))
}

func Sha1(raw string) string {
	h := sha1.New()
	io.WriteString(h, raw)

	return fmt.Sprintf("%x", h.Sum(nil))
}
