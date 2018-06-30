package xstring

import (
	"bytes"
	"unicode"
)

func IsChineseChar(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func Concat(pieces ...string) string {
	var buffer bytes.Buffer
	for _, piece := range pieces {
		buffer.WriteString(piece)
	}
	return buffer.String()
}
