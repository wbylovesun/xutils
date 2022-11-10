package xslice

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func RmDuplicate[T SliceElementType](data []T) []T {
	var nData []T
	var nMap = map[T]bool{}
	for _, v := range data {
		if _, ok := nMap[v]; !ok {
			nData = append(nData, v)
			nMap[v] = true
		}
	}
	return nData
}

func Sort[T SliceElementType](data []T) []T {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	return data
}

func BisectionSearch[T SliceElementType](list []T, target T) int {
	Sort(list)
	low := 0
	high := len(list) - 1
	for {
		if low > high {
			return -1
		}
		mid := (high + low) / 2
		guess := list[mid]
		if guess == target {
			return mid
		} else if guess > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}

func Filter[T SliceElementType](data []T, f func(v T, k int) bool) []T {
	var nData []T
	for k, v := range data {
		if f(v, k) {
			continue
		}
		nData = append(nData, v)
	}
	return nData
}

// Contains 判断数组和切片中是否存在某个值
func Contains[T SliceElementType](data []T, ct T) bool {
	if data == nil {
		return false
	}
	for _, v := range data {
		if v == ct {
			return true
		}
	}
	return false
}

func ContainsSlice[T SliceElementType](data []T, ct []T) bool {
	nData := RmDuplicate(data)
	nCt := RmDuplicate(ct)
	Sort(nData)
	Sort(nCt)

	i := -1
	for _, w := range nCt {
		nData = nData[i+1:]
		mid := BisectionSearch(nData, w)
		if mid == -1 {
			return false
		}
		i = mid
	}
	return true
}

func Walk[T SliceElementType](data []T, f func(v T, k int) T) {
	for k, v := range data {
		data[k] = f(v, k)
	}
}

func Map[T SliceElementType](f func(v T) T, data ...[]T) []T {
	nData := make([]T, 0, len(data))
	for _, v := range data {
		for _, vv := range v {
			nData = append(nData, f(vv))
		}
	}
	return nData
}

func Min[T SliceElementType](data []T) *T {
	if len(data) == 0 {
		return nil
	}
	var min T = data[0]
	for _, v := range data {
		if min > v {
			min = v
		}
	}
	return &min
}

func Max[T SliceElementType](data []T) *T {
	if len(data) == 0 {
		return nil
	}
	var max T = data[0]
	for _, v := range data {
		if max < v {
			max = v
		}
	}
	return &max
}

func Join[T SliceElementType](data []T, sep string) string {
	s := make([]string, len(data))
	var typeDetector T
	t := reflect.TypeOf(typeDetector)
	if t.Kind() == reflect.String {
		for k, v := range data {
			s[k] = reflect.ValueOf(v).String()
		}
	} else {
		for k, v := range data {
			s[k] = fmt.Sprintf("%d", v)
		}
	}
	return strings.Join(s, sep)
}
