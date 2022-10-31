package xcalc

import (
	"github.com/wbylovesun/xutils"
)

// SequenceValuePad 当前不存在的段以前一个段数值填充，否则保留当前段的值
func SequenceValuePad[T any](sequence []string, m map[string]T, f func(last T, k int) T) map[string]T {
	if f == nil {
		return nil
	}

	copied := xutils.Copy(m)
	mx := copied.(map[string]T)

	var last T
	for k, v := range sequence {
		mv, ok := mx[v]
		if !ok {
			mx[v] = f(last, k)
		} else {
			last = mv
		}
	}
	return mx
}

// IncrementalSequenceValueDiff 给递增的时间值序列计算差值
//
// 如sequence("a", "b", "c", "d", "e"), m("a": 1, "d": 3)，则结果为m("a":1,"b":0,"c":0,"d":2,"e":0)
//
// f不给定时，返回空切片
//
// p不给定时，需要在外部对m进行补齐至与sequence一致，即m所有的key与sequence的值是一致的
func IncrementalSequenceValueDiff[T any](sequence []string, m map[string]T, f func(current, last T) T, p func(v T, k int) T) map[string]T {
	if len(sequence) == 0 {
		return nil
	}
	if f == nil {
		return nil
	}
	if p == nil && len(sequence) != len(m) {
		return nil
	}

	var mx map[string]T
	if p != nil {
		mx = SequenceValuePad(sequence, m, p)
	} else {
		copied := xutils.Copy(m)
		mx = copied.(map[string]T)
	}
	var last T
	for k, v := range sequence {
		if k == 0 {
			last = mx[v]
			continue
		}
		last, mx[v] = mx[v], f(mx[v], last)
	}
	return mx
}
