package xstrings

import (
	"github.com/wbylovesun/xutils/types"
	"reflect"
	"strconv"
	"strings"
)

func ToNumberSlice[T types.Number](str string, separation string, to T) []T {
	splits := strings.Split(str, separation)
	var slice = make([]T, 0, len(splits))
	for _, split := range splits {
		switch reflect.TypeOf(to).Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			fallthrough
		case reflect.Uint:
			fallthrough
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint64:
			slice = convertNumber(slice, split)
		case reflect.Float32:
			slice = convertFloat(slice, split, 32)
		case reflect.Float64:
			slice = convertFloat(slice, split, 64)
		}
	}
	return slice
}

func convertNumber[T types.Number](slice []T, v string) []T {
	v = strings.Trim(v, " \r\n\t\b")
	atoi, err := strconv.Atoi(v)
	if err != nil {
		return slice
	}
	slice = append(slice, T(atoi))
	return slice
}

func convertFloat[T types.Number](slice []T, v string, bitsize int) []T {
	f, err := strconv.ParseFloat(v, bitsize)
	if err != nil {
		return slice
	}
	slice = append(slice, T(f))
	return slice
}

func ToIntSlice(str string, separation string) []int {
	return ToNumberSlice(str, separation, int(0))
}

func ToFloat64Slice(str string, separation string) []float64 {
	return ToNumberSlice(str, separation, float64(0))
}

func ToStringSlice(str string, separation string) []string {
	splits := strings.Split(str, separation)
	return splits
}
