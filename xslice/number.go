package xslice

import (
	"reflect"
	"strconv"
	"strings"
)

// Join 将一个slice使用sep连接成一个字符串
// slice的元素当为float类型时，当小数为0时会被舍弃；小数不为0时，保留17-整数位长度的小数
func Join[T NumberSliceElementType | UnsignedNumberSliceElementType | float32 | float64](t []T, sep string) string {
	if len(t) == 0 {
		return ""
	}
	var s strings.Builder
	var l = len(t)
	typ := reflect.ValueOf(t[0]).Kind()
	switch typ {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		for i, v := range t {
			s.WriteString(strconv.FormatInt(int64(v), 10))
			if i != l-1 {
				s.WriteString(sep)
			}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		for i, v := range t {
			s.WriteString(strconv.FormatUint(uint64(v), 10))
			if i != l-1 {
				s.WriteString(sep)
			}
		}
	case reflect.Float32, reflect.Float64:
		for i, v := range t {
			s.WriteString(strconv.FormatFloat(float64(v), 'f', -1, 64))
			if i != l-1 {
				s.WriteString(sep)
			}
		}
	default:
		return ""
	}
	return s.String()
}

func ToInt64Slice[T NumberSliceElementType | UnsignedNumberSliceElementType](t []T) []int64 {
	return ToNumberSlice(t, int64(0))
}

func ToInt32Slice[T NumberSliceElementType | UnsignedNumberSliceElementType](t []T) []int32 {
	return ToNumberSlice(t, int32(0))
}

func ToInt16Slice[T NumberSliceElementType | UnsignedNumberSliceElementType](t []T) []int16 {
	return ToNumberSlice(t, int16(0))
}

func ToInt8Slice[T NumberSliceElementType | UnsignedNumberSliceElementType](t []T) []int8 {
	return ToNumberSlice(t, int8(0))
}

func ToIntSlice[T NumberSliceElementType | UnsignedNumberSliceElementType](t []T) []int {
	return ToNumberSlice(t, int(0))
}

func ToNumberSlice[T NumberSliceElementType | UnsignedNumberSliceElementType, U NumberSliceElementType](t []T, u U) []U {
	if len(t) == 0 {
		return nil
	}
	var us = make([]U, len(t))
	for i, v := range t {
		us[i] = U(v)
	}
	return us
}
