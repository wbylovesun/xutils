package xvalidator

import (
	"reflect"
	"strconv"
	"time"
)

// asInt returns the parameter as a int64
// or panics if it can't convert
func asInt(param string) int64 {
	i, err := strconv.ParseInt(param, 0, 64)
	panicIf(err)

	return i
}

// asIntFromTimeDuration parses param as time.Duration and returns it as int64
// or panics on error.
func asIntFromTimeDuration(param string) int64 {
	d, err := time.ParseDuration(param)
	if err != nil {
		// attempt parsing as an an integer assuming nanosecond precision
		return asInt(param)
	}
	return int64(d)
}

// asIntFromType calls the proper function to parse param as int64,
// given a field's Type t.
func asIntFromType(t reflect.Type, param string) int64 {
	switch t {
	case timeDurationType:
		return asIntFromTimeDuration(param)
	default:
		return asInt(param)
	}
}

// asUint returns the parameter as a uint64
// or panics if it can't convert
func asUint(param string) uint64 {

	i, err := strconv.ParseUint(param, 0, 64)
	panicIf(err)

	return i
}

// asFloat64 returns the parameter as a float64
// or panics if it can't convert
func asFloat64(param string) float64 {
	i, err := strconv.ParseFloat(param, 64)
	panicIf(err)
	return i
}

// asFloat32 returns the parameter as a float32
// or panics if it can't convert
func asFloat32(param string) float64 {
	i, err := strconv.ParseFloat(param, 32)
	panicIf(err)
	return i
}

// asBool returns the parameter as a bool
// or panics if it can't convert
func asBool(param string) bool {

	i, err := strconv.ParseBool(param)
	panicIf(err)

	return i
}

func panicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func isEmbedded(struct1 interface{}, struct2 interface{}) bool {
	struct1Type := reflect.TypeOf(struct1)
	struct2Type := reflect.TypeOf(struct2)

	// 遍历 struct2 的所有字段，检查是否有内嵌字段与 struct1 类型匹配
	for i := 0; i < struct2Type.NumField(); i++ {
		field := struct2Type.Field(i)
		if field.Anonymous && field.Type == struct1Type {
			return true
		}
	}

	return false
}

// getEmbeddedStruct 查找并返回内嵌的结构体（如果存在）
func getEmbeddedStruct(b interface{}, embeddedType interface{}) (interface{}, bool) {
	bValue := reflect.ValueOf(b)
	bType := reflect.TypeOf(b)
	embeddedTypeValue := reflect.TypeOf(embeddedType)

	// 遍历 B 的所有字段
	for i := 0; i < bType.NumField(); i++ {
		field := bType.Field(i)
		// 判断该字段是否是匿名字段，且类型与 embeddedType 匹配
		if field.Anonymous && field.Type == embeddedTypeValue {
			// 获取该字段的值
			return bValue.Field(i).Interface(), true
		}
	}

	return nil, false
}
