package xslice

import (
	"reflect"
	"strconv"
	"strings"
)

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
	case reflect.Float32 | reflect.Float64:
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
	if len(t) == 0 {
		return nil
	}
	var i64 = make([]int64, len(t))
	for i, v := range t {
		i64[i] = int64(v)
	}
	return i64
}

func ToInt32Slice[T NumberSliceElementType | UnsignedNumberSliceElementType](t []T) []int32 {
	if len(t) == 0 {
		return nil
	}
	var i32 = make([]int32, len(t))
	for i, v := range t {
		i32[i] = int32(v)
	}
	return i32
}

func ToInt16Slice[T NumberSliceElementType | UnsignedNumberSliceElementType](t []T) []int16 {
	if len(t) == 0 {
		return nil
	}
	var i16 = make([]int16, len(t))
	for i, v := range t {
		i16[i] = int16(v)
	}
	return i16
}

func ToInt8Slice[T NumberSliceElementType | UnsignedNumberSliceElementType](t []T) []int8 {
	if len(t) == 0 {
		return nil
	}
	var i8 = make([]int8, len(t))
	for i, v := range t {
		i8[i] = int8(v)
	}
	return i8
}

func ToIntSlice[T NumberSliceElementType | UnsignedNumberSliceElementType](t []T) []int {
	if len(t) == 0 {
		return nil
	}
	var is = make([]int, len(t))
	for i, v := range t {
		is[i] = int(v)
	}
	return is
}
