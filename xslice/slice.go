package xslice

import (
	"sort"
)

func Clone[T SliceElementType](data []T) []T {
	t := make([]T, len(data))
	copy(t, data)
	return t
}

func Sort[T SliceElementType](data []T) {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
}

func CloneSort[T SliceElementType](data []T) []T {
	t := Clone(data)
	Sort(t)
	return t
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
	f := map[T]struct{}{}
	var t []T
	for _, ele := range elements {
		if _, ok := f[ele]; !ok {
			f[ele] = struct{}{}
			t = append(t, ele)
		}
	}
	return t
}

func Distinct[T SliceElementType](elements []T) []T {
	return Unique(elements)
}

func Intersection[T SliceElementType](A, B []T) []T {
	a := A
	b := B
	if len(a) > len(b) {
		a, b = b, a
	}
	var inter []T
	for _, i := range a {
		if Contains(b, i) {
			inter = append(inter, i)
		}
	}
	return inter
}

func CloneDiff[T SliceElementType](A, B []T) []T {
	a := CloneSort(A)
	b := CloneSort(B)
	return diff(a, b)
}

func Diff[T SliceElementType](A, B []T) []T {
	Sort(A)
	Sort(B)
	return diff(A, B)
}

func diff[T SliceElementType](A, B []T) []T {
	var diffed []T
	var t = make(map[T]struct{})
	for _, i := range B {
		t[i] = struct{}{}
	}
	for _, i := range A {
		if _, ok := t[i]; !ok {
			diffed = append(diffed, i)
		}
	}
	return diffed
}

func Union[T SliceElementType](A, B []T) []T {
	return Unique(append(A, B...))
}
