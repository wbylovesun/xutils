package xstrings

import (
	"strconv"
	"strings"
)

func ToIntSlice(str string, separation string) []int {
	splits := strings.Split(str, separation)
	var slice = make([]int, 0, len(splits))
	for _, split := range splits {
		atoi, err := strconv.Atoi(split)
		if err != nil {
			continue
		}
		slice = append(slice, atoi)
	}
	return slice
}

func ToStringSlice(str string, separation string) []string {
	splits := strings.Split(str, separation)
	return splits
}
