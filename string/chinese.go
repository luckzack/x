package utils

import (
	"regexp"
	"unicode"
)

var ContainsChinese = containsChineseByUnicode

func containsChineseByUnicode(str string) bool {

	for _, c := range str {
		if unicode.Is(unicode.Han, c) {
			return true
		}
	}
	return false
}

// TODO：这个还不能识别üe这样的拼音
func containsChineseByRegexp(str string) bool {
	for _, c := range str {
		if regexp.MustCompile("[\u3002\uff1b\uff0c\uff1a\u201c\u201d\uff08\uff09\u3001\uff1f\u300a\u300b]").MatchString(string(c)) {
			return false
		}
	}
	return true
}
