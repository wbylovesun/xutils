package main

import (
	"fmt"
	"github.com/wbylovesun/xutils/xslice"
)

func main() {
	a := []int{8, 10, 13, 22, 19, 7, 6, 4, 9, 1, 5, 2, 11, 17, 12, 16, 15, 21, 14, 18, 20, 3}
	b := []int{20, 18, 21, 5, 17, 3, 4, 6, 9, 1, 2, 11, 16, 21, 14, 15, 8, 13, 10}
	//
	fmt.Println(a, b)

	// 求交集
	fmt.Println(xslice.Intersection(a, b))
	fmt.Println(a, b)

	// 保持原切片，克隆后求差集
	fmt.Println(xslice.CloneDiff(a, b))
	fmt.Println(a, b)

	// 直接求差集
	c := xslice.Clone(a)
	d := xslice.Clone(b)
	fmt.Println(xslice.Diff(c, d))
	fmt.Println(c, d)

	// 求并集
	fmt.Println(xslice.Union(a, b))
	fmt.Println(a, b)

	// 去重
	fmt.Println(xslice.Distinct(a))

	// 拼字符串
	fmt.Println(xslice.Join([]int{1, 2, 3, 4, 5}, "~"))
	fmt.Println(xslice.Join([]float64{11118.0122229292929304958159, -20.0, 100.99, 80.88}, ","))

	// 求最小/大值
	fmt.Println(a, *xslice.Min(a))
	fmt.Println(b, *xslice.Max(b))

	// 转其它整型类型的切片
	fmt.Println(xslice.ToInt64Slice([]int{1, 2, 3, 4, 5}))
	fmt.Println(xslice.ToInt32Slice([]int{1, 2, 3, 4, 5}))
	fmt.Println(xslice.ToInt16Slice([]int{1, 2, 3, 4, 5}))
	fmt.Println(xslice.ToInt8Slice([]int{1, 2, 3, 4, 5}))
	fmt.Println(xslice.ToIntSlice([]int{1, 2, 3, 4, 5}))
}
