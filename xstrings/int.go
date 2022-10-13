package xstrings

import "strconv"

func MustInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(`MustInt(` + s + `): ` + err.Error())
	}
	return i
}

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
