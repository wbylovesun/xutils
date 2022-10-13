package xstrings

import (
	"fmt"
	"strings"
)

// Lcfirst 如果第一个字符在A-Z之间，则将其变为小写
func Lcfirst(s string) string {
	if len(s) == 0 {
		return s
	}
	c := s[0]
	if c >= 'A' && c <= 'Z' {
		first := strings.ToLower(string(c))
		if len(s) > 1 {
			return first + s[1:]
		}
		return first
	}
	return s
}

// Ucfirst 如果第一个字符在a-z之间，则将其变为大写
func Ucfirst(s string) string {
	if len(s) == 0 {
		return s
	}
	c := s[0]
	if c >= 'a' && c <= 'z' {
		first := strings.ToUpper(string(c))
		if len(s) > 1 {
			return first + s[1:]
		}
		return first
	}
	return s
}

// StartsWith 字符串s是否以p开头
func StartsWith(s, p string) bool {
	sl := len(s)
	if sl == 0 {
		return false
	}
	pl := len(p)
	if pl == 0 {
		return false
	}
	if sl < pl {
		return false
	}
	return strings.Compare(s[0:pl], p) == 0
}

// EndsWith 字符串s是否以p结尾
func EndsWith(s, p string) bool {
	sl := len(s)
	if sl == 0 {
		return false
	}
	pl := len(p)
	if pl == 0 {
		return false
	}
	if sl < pl {
		return false
	}
	sp := sl - pl
	fmt.Println(sp, sl, s[sp:sl])
	return strings.Compare(s[sp:sl], p) == 0
}
