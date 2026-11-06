package xslice

import (
	"sort"
)

// Clone 返回切片的浅拷贝（新底层数组），修改副本不影响原切片；nil 输入返回空切片。
func Clone[T any](data []T) []T {
	t := make([]T, len(data))
	copy(t, data)
	return t
}

// Sort 对切片原地升序排序，会修改传入的切片。
func Sort[T SliceElementType](data []T) {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
}

// CloneSort 返回升序排序后的新副本，原切片保持不变。
func CloneSort[T SliceElementType](data []T) []T {
	t := Clone(data)
	Sort(t)
	return t
}

// BisectionSearch 二分查找 target，返回其在排序后切片中的下标，未找到时返回 -1。
// 注意：查找前会先对 list 原地排序（会修改调用方传入的切片），返回的下标基于排序后的顺序。
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

// Filter 剔除切片中谓词返回 true 的元素（f(v, k) 返回 true 时该元素被剔除），返回剩余元素组成的新切片。
func Filter[T any](data []T, f func(v T, k int) bool) []T {
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
func Contains[T comparable](data []T, ct T) bool {
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

// ContainsSlice 判断 ct 中的每个元素是否都存在于 data 中。
// 双方会先去重再比较，按集合语义处理，不考虑元素的重复次数。
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

// Walk 对切片的每个元素执行 f 并原地写回（data[k] = f(v, k)），会修改传入的切片。
func Walk[T any](data []T, f func(v T, k int) T) {
	for k, v := range data {
		data[k] = f(v, k)
	}
}

// Map 对一个或多个切片中的每个元素依次执行 f，结果按传入顺序拼接为新的切片返回。
func Map[T any](f func(v T) T, data ...[]T) []T {
	nData := make([]T, 0, len(data))
	for _, v := range data {
		for _, vv := range v {
			nData = append(nData, f(vv))
		}
	}
	return nData
}

// Min 返回切片中最小元素的指针，空切片返回 nil。
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

// Max 返回切片中最大元素的指针，空切片返回 nil。
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

// Unique 去除重复元素，保留首次出现的顺序，返回新切片。
func Unique[T comparable](elements []T) []T {
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

// Distinct 去除重复元素，保留首次出现的顺序，返回新切片（等同 Unique）。
func Distinct[T comparable](elements []T) []T {
	return Unique(elements)
}

// Intersection 返回 A 和 B 的交集；以较短的切片为遍历基准，结果不去重。
func Intersection[T comparable](A, B []T) []T {
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

// CloneDiff 返回 A 中不在 B 里的元素；先复制再排序，不修改入参切片。
func CloneDiff[T SliceElementType](A, B []T) []T {
	a := CloneSort(A)
	b := CloneSort(B)
	return diff(a, b)
}

// Diff 返回 A 中不在 B 里的元素。注意：会原地排序 A 和 B，不希望修改入参时请使用 CloneDiff。
func Diff[T SliceElementType](A, B []T) []T {
	Sort(A)
	Sort(B)
	return diff(A, B)
}

// diff 返回 A 中不在 B 里的元素，保留 A 侧的顺序与重复元素。
func diff[T comparable](A, B []T) []T {
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

// Union 返回 A 与 B 合并去重后的新切片，保留首次出现的顺序。
func Union[T comparable](A, B []T) []T {
	return Unique(append(A, B...))
}

// ConvertTo 将 from 中的每个元素经 f 转换后组成新切片返回。
func ConvertTo[T any, O any](from []T, f func(T) O) []O {
	t := make([]O, 0, len(from))
	for _, v := range from {
		t = append(t, f(v))
	}
	return t
}

// PluckToMap 将每个元素经 f 拆解为键值对后构建 map 返回，键冲突时后写入者覆盖。
func PluckToMap[T any, K comparable, V any](from []T, f func(T) (K, V)) map[K]V {
	m := make(map[K]V, 0)
	for _, item := range from {
		k, v := f(item)
		m[k] = v
	}
	return m
}

// PtrSlice 将值切片转换为指针切片；每个指针指向元素的独立副本，修改其指向的值不影响原切片。
func PtrSlice[V any](vals []V) []*V {
	out := make([]*V, len(vals))
	for i, v := range vals {
		vv := v
		out[i] = &vv
	}
	return out
}

// ValueSlice 将指针切片转换为值切片，nil 元素取零值，返回与入参不共享内存的新切片。
func ValueSlice[V any](vals []*V) []V {
	out := make([]V, len(vals))
	for i, v := range vals {
		if v != nil {
			out[i] = *v
		}
	}
	return out
}
