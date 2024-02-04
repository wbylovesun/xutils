package xslice

import (
	"sort"
)

// RmDuplicate Deprecated 移除重复项
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
	nData := Unique(data)
	nCt := Unique(ct)
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

func Unique[T SliceElementType](elements []T) []T {
	f := map[T]bool{}
	var t []T
	for _, ele := range elements {
		if _, ok := f[ele]; !ok {
			f[ele] = true
			t = append(t, ele)
		}
	}
	return t
}

func Diff[T SliceElementType](A, B []T) []T {
	var diffed []T
	for _, i := range A {
		if !Contains(B, i) {
			diffed = append(diffed, i)
		}
	}
	return diffed
}
