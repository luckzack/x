package validator

import (
	"fmt"
	"regexp"
	"strings"
)

func IsMatch(val, pattern string) bool {
	match, err := regexp.Match(pattern, []byte(val))
	if err != nil {
		return false
	}

	return match
}

func IsEnglishIdentifier(val string, pattern ...string) bool {
	defpattern := "^[a-zA-Z0-9\\-\\_\\.]+$"
	if len(pattern) > 0 {
		defpattern = pattern[0]
	}

	return IsMatch(val, defpattern)
}

func IsMail(val string) bool {
	return IsMatch(val, `\w[-._\w]*@\w[-._\w]*\.\w+`)
}

func IsPhone(val string) bool {
	if strings.HasPrefix(val, "+") {
		return IsMatch(val[1:], `\d{13}`)
	} else {
		return IsMatch(val, `\d{11}`)
	}
}

func IsIPStr(s string) bool {
	reg := regexp.MustCompile(`^(?:(?:[01]?\d{1,2}|2[0-4]\d|25[0-5])\.){3}(?:[01]?\d{1,2}|2[0-4]\d|25[0-5])$`)
	return reg.MatchString(s)
}

//四个字节，32位整数，int32
func IsIPInt(i int) bool {
	if i > 4294967295 || i < 0 {
		return false
	}
	return true
}

func IsURL(s string) bool {
	reg := regexp.MustCompile(`[a-zA-z]+://[^\s]*`)
	return reg.MatchString(s)
}

//scheme:http,rtmp,rtmfp
func IsSpecifiedURL(s, scheme string) bool {
	reg := regexp.MustCompile(scheme + `://[^\s]*`)
	return reg.MatchString(s)
}

//Unix directory starts by `/`
func IsUnixDir(s string) bool {
	return strings.HasPrefix(s, `/`)
}

//只能由字母、数字和下划线组合，且首字符不能是数字
func IsIllegalName(s string) bool {
	b := []byte(s)
	if b[0] >= 48 && b[0] <= 57 {
		return false
	}
	reg := regexp.MustCompile(`^[A-Za-z0-9_]+$`)
	return reg.MatchString(s)
}

//10个字符长度时间戳，s单位，以1开头
//如果是数字，必须先转化为int64
func IsTimestamp(v interface{}) bool {
	reg := regexp.MustCompile(`^1\d{9}$`)

	if vv, ok := v.(string); ok {
		return reg.MatchString(vv)
	} else if vv, ok := v.(int64); ok {
		return vv > 1000000000 && vv < 1999999999
	}
	fmt.Println("!!!")
	return false
}
