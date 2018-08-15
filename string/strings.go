package string

import (
	"bytes"
	"strings"
)

func ConcatLong(pieces ...string) string {
	var buffer bytes.Buffer
	for _, piece := range pieces {
		buffer.WriteString(piece)
	}
	return buffer.String()
}

func ConcatShort(pieces ...string) string {
	return strings.Join(pieces, "")
}
