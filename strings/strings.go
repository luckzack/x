package xstring

import (
	"bytes"
)

func Concat(pieces ...string) string {
	var buffer bytes.Buffer
	for _, piece := range pieces {
		buffer.WriteString(piece)
	}
	return buffer.String()
}
