package sha1

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"gitee.com/johng/gf/g/util/gconv"
)

// 将任意类型的变量进行SHA摘要(注意map等非排序变量造成的不同结果)
// 内部使用了md5计算，因此效率会稍微差一些，更多情况请使用 EncodeString
func Encrypt(v interface{}) string {
	r := sha1.Sum(gconv.Bytes(v))
	return hex.EncodeToString(r[:])
}

// 对字符串行SHA1摘要计算
func EncryptString(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

// 对文件内容进行SHA1摘要计算
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
