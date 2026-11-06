package xslice

import (
	"reflect"
	"strconv"
	"testing"
)

func assertEqual[T any](t *testing.T, name string, got, want []T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s = %v, want %v", name, got, want)
	}
}

func TestClone(t *testing.T) {
	src := []int{1, 2, 3}
	got := Clone(src)
	assertEqual(t, "Clone(src)", got, []int{1, 2, 3})

	// 修改副本不应影响原切片（不共享底层数组）
	got[0] = 100
	if src[0] != 1 {
		t.Errorf("Clone 与原切片共享底层数组：修改副本后 src[0] = %d", src[0])
	}

	if n := len(Clone([]int{})); n != 0 {
		t.Errorf("len(Clone([]int{})) = %d, want 0", n)
	}
	if n := len(Clone[int](nil)); n != 0 {
		t.Errorf("len(Clone[int](nil)) = %d, want 0", n)
	}
}

func TestSort(t *testing.T) {
	ints := []int{3, 1, 2, 5, 4}
	Sort(ints)
	assertEqual(t, "Sort(ints)", ints, []int{1, 2, 3, 4, 5})

	strs := []string{"b", "a", "c"}
	Sort(strs)
	assertEqual(t, "Sort(strs)", strs, []string{"a", "b", "c"})

	empty := []int{}
	Sort(empty)
	assertEqual(t, "Sort(empty)", empty, []int{})
}

func TestCloneSort(t *testing.T) {
	src := []int{3, 1, 2}
	got := CloneSort(src)
	assertEqual(t, "CloneSort(src)", got, []int{1, 2, 3})
	assertEqual(t, "CloneSort 后原切片", src, []int{3, 1, 2}) // 不应被修改
}

func TestBisectionSearch(t *testing.T) {
	list := []int{5, 1, 4, 2, 3}
	if idx := BisectionSearch(list, 3); idx < 0 || list[idx] != 3 {
		t.Errorf("BisectionSearch(list, 3) = %d, 应命中元素 3", idx)
	}
	if idx := BisectionSearch(list, 10); idx != -1 {
		t.Errorf("BisectionSearch(list, 10) = %d, want -1", idx)
	}
	if idx := BisectionSearch([]int{}, 1); idx != -1 {
		t.Errorf("BisectionSearch([], 1) = %d, want -1", idx)
	}
	if idx := BisectionSearch[int](nil, 1); idx != -1 {
		t.Errorf("BisectionSearch(nil, 1) = %d, want -1", idx)
	}

	// 重复元素：应命中任意一个相等的元素
	dup := []int{1, 2, 2, 3}
	if idx := BisectionSearch(dup, 2); idx < 1 || idx > 2 {
		t.Errorf("BisectionSearch(dup, 2) = %d, 应命中下标 1 或 2", idx)
	}

	if idx := BisectionSearch([]string{"b", "a"}, "a"); idx < 0 {
		t.Errorf("BisectionSearch(string) = %d, 应命中元素", idx)
	}

	// 文档化副作用：查找会对输入切片原地排序
	mut := []int{2, 1}
	_ = BisectionSearch(mut, 2)
	assertEqual(t, "BisectionSearch 原地排序输入", mut, []int{1, 2})
}

func TestFilter(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	// Filter 为剔除语义：f 返回 true 的元素会被剔除
	got := Filter(data, func(v, k int) bool { return v%2 == 0 })
	assertEqual(t, "Filter(剔除偶数)", got, []int{1, 3, 5})

	if n := len(Filter([]int{}, func(v, k int) bool { return true })); n != 0 {
		t.Errorf("len(Filter([], ...)) = %d, want 0", n)
	}
}

func TestContains(t *testing.T) {
	if !Contains([]int{1, 2, 3}, 2) {
		t.Error("Contains([1,2,3], 2) = false, want true")
	}
	if Contains([]int{1, 2, 3}, 9) {
		t.Error("Contains([1,2,3], 9) = true, want false")
	}
	if Contains([]int{}, 1) {
		t.Error("Contains([], 1) = true, want false")
	}
	if Contains[int](nil, 1) {
		t.Error("Contains(nil, 1) = true, want false")
	}
	if !Contains([]string{"a", "b"}, "b") {
		t.Error(`Contains(["a","b"], "b") = false, want true`)
	}
}

func TestContainsSlice(t *testing.T) {
	if !ContainsSlice([]int{3, 1, 2}, []int{2, 1}) {
		t.Error("ContainsSlice([3,1,2], [2,1]) = false, want true")
	}
	if ContainsSlice([]int{1, 2}, []int{3}) {
		t.Error("ContainsSlice([1,2], [3]) = true, want false")
	}
	if ContainsSlice([]int{1}, []int{1, 2}) {
		t.Error("ContainsSlice([1], [1,2]) = true, want false")
	}
	if !ContainsSlice([]int{1, 2}, []int{}) {
		t.Error("ContainsSlice([1,2], []) = false, want true（空 ct 视为包含）")
	}
	if !ContainsSlice[int](nil, nil) {
		t.Error("ContainsSlice(nil, nil) = false, want true")
	}
	if ContainsSlice([]int{}, []int{1}) {
		t.Error("ContainsSlice([], [1]) = true, want false")
	}
	// 注意：实现先对 ct 做 Unique 去重，因此重复元素按集合语义处理。
	// 若期望多重集语义（数据中必须有足够数量的重复元素），此处应为 false，见问题清单。
	if !ContainsSlice([]int{1}, []int{1, 1}) {
		t.Error("ContainsSlice([1], [1,1]) = false, want true（按集合语义）")
	}
}

func TestWalk(t *testing.T) {
	data := []int{1, 2, 3}
	Walk(data, func(v, k int) int { return v * 2 })
	assertEqual(t, "Walk 后", data, []int{2, 4, 6})

	lastK := -1
	Walk(data, func(v, k int) int { lastK = k; return v })
	if lastK != 2 {
		t.Errorf("Walk 回调最后收到的 k = %d, want 2", lastK)
	}
}

func TestMap(t *testing.T) {
	got := Map(func(v int) int { return v * 2 }, []int{1, 2}, []int{3})
	assertEqual(t, "Map 多切片", got, []int{2, 4, 6})

	if n := len(Map(func(v int) int { return v })); n != 0 {
		t.Errorf("len(Map(无切片)) = %d, want 0", n)
	}
}

func TestMin(t *testing.T) {
	if got := Min([]int{3, 1, 2}); got == nil || *got != 1 {
		t.Errorf("Min([3,1,2]) = %v, want 指向 1", got)
	}
	if got := Min([]int{-1, -5}); got == nil || *got != -5 {
		t.Errorf("Min([-1,-5]) = %v, want 指向 -5", got)
	}
	if got := Min([]int{}); got != nil {
		t.Errorf("Min([]) = %v, want nil", got)
	}
	if got := Min([]int{7}); got == nil || *got != 7 {
		t.Errorf("Min([7]) = %v, want 指向 7", got)
	}
	if got := Min([]string{"b", "a"}); got == nil || *got != "a" {
		t.Errorf(`Min(["b","a"]) = %v, want 指向 "a"`, got)
	}
}

func TestMax(t *testing.T) {
	if got := Max([]int{3, 1, 2}); got == nil || *got != 3 {
		t.Errorf("Max([3,1,2]) = %v, want 指向 3", got)
	}
	if got := Max([]int{}); got != nil {
		t.Errorf("Max([]) = %v, want nil", got)
	}
	if got := Max([]int{-7}); got == nil || *got != -7 {
		t.Errorf("Max([-7]) = %v, want 指向 -7", got)
	}
}

func TestUnique(t *testing.T) {
	assertEqual(t, "Unique", Unique([]int{1, 2, 1, 3, 2}), []int{1, 2, 3})
	assertEqual(t, "Unique strings", Unique([]string{"a", "b", "a"}), []string{"a", "b"})
	if n := len(Unique([]int{})); n != 0 {
		t.Errorf("len(Unique([])) = %d, want 0", n)
	}
	if n := len(Unique[int](nil)); n != 0 {
		t.Errorf("len(Unique[int](nil)) = %d, want 0", n)
	}
}

func TestDistinct(t *testing.T) {
	assertEqual(t, "Distinct", Distinct([]int{1, 1, 2}), []int{1, 2})
}

func TestIntersection(t *testing.T) {
	assertEqual(t, "Intersection", Intersection([]int{1, 2, 3}, []int{2, 3, 4}), []int{2, 3})

	if n := len(Intersection([]int{1, 2}, []int{3, 4})); n != 0 {
		t.Errorf("len(Intersection(无交集)) = %d, want 0", n)
	}
	if n := len(Intersection([]int{}, []int{1})); n != 0 {
		t.Errorf("len(Intersection([], [1])) = %d, want 0", n)
	}
	if n := len(Intersection([]int{1}, []int{})); n != 0 {
		t.Errorf("len(Intersection([1], [])) = %d, want 0", n)
	}
}

func TestDiff(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := []int{2, 4}
	got := Diff(a, b)
	assertEqual(t, "Diff(a, b)", got, []int{1, 3})
	// 注意：Diff 会原地排序 a 和 b（现有行为，见问题清单）
}

func TestCloneDiff(t *testing.T) {
	a := []int{3, 1, 2}
	b := []int{2}
	got := CloneDiff(a, b)
	assertEqual(t, "CloneDiff(a, b)", got, []int{1, 3})
	assertEqual(t, "CloneDiff 后原 a", a, []int{3, 1, 2}) // 不应被修改
	assertEqual(t, "CloneDiff 后原 b", b, []int{2})       // 不应被修改
}

func TestUnion(t *testing.T) {
	assertEqual(t, "Union", Union([]int{1, 2}, []int{2, 3}), []int{1, 2, 3})
	assertEqual(t, "Union nil A", Union[int](nil, []int{1}), []int{1})
}

func TestConvertTo(t *testing.T) {
	got := ConvertTo([]int{1, 2, 3}, func(v int) string { return strconv.Itoa(v) })
	assertEqual(t, "ConvertTo", got, []string{"1", "2", "3"})
}

func TestPluckToMap(t *testing.T) {
	type kv struct {
		K int
		V string
	}
	toPair := func(item kv) (int, string) { return item.K, item.V }

	got := PluckToMap([]kv{{1, "a"}, {2, "b"}}, toPair)
	if want := map[int]string{1: "a", 2: "b"}; !reflect.DeepEqual(got, want) {
		t.Errorf("PluckToMap = %v, want %v", got, want)
	}

	// 键冲突时后写入者覆盖
	got2 := PluckToMap([]kv{{1, "a"}, {1, "c"}}, toPair)
	if want2 := map[int]string{1: "c"}; !reflect.DeepEqual(got2, want2) {
		t.Errorf("PluckToMap(键冲突) = %v, want %v", got2, want2)
	}

	if n := len(PluckToMap([]kv(nil), toPair)); n != 0 {
		t.Errorf("len(PluckToMap(nil)) = %d, want 0", n)
	}
}

func TestPtrSlice(t *testing.T) {
	src := []int{1, 2, 3}
	got := PtrSlice(src)
	if len(got) != 3 || *got[0] != 1 || *got[1] != 2 || *got[2] != 3 {
		t.Fatalf("PtrSlice(%v) = %v, want 每个指针指向对应元素", src, got)
	}

	// 修改指针指向的值不应影响原切片（元素为独立副本）
	*got[0] = 100
	if src[0] != 1 {
		t.Errorf("PtrSlice 与原切片共享元素：修改后 src[0] = %d", src[0])
	}

	if n := len(PtrSlice([]int{})); n != 0 {
		t.Errorf("len(PtrSlice([])) = %d, want 0", n)
	}
	if n := len(PtrSlice[int](nil)); n != 0 {
		t.Errorf("len(PtrSlice[int](nil)) = %d, want 0", n)
	}
}

func TestValueSlice(t *testing.T) {
	a, b := 1, 2
	got := ValueSlice([]*int{&a, nil, &b})
	if len(got) != 3 || got[0] != 1 || got[1] != 0 || got[2] != 2 {
		t.Errorf("ValueSlice([%d, nil, %d]) = %v, want [1 0 2]", a, b, got)
	}

	if n := len(ValueSlice([]*int{})); n != 0 {
		t.Errorf("len(ValueSlice([])) = %d, want 0", n)
	}
	if n := len(ValueSlice[int](nil)); n != 0 {
		t.Errorf("len(ValueSlice[int](nil)) = %d, want 0", n)
	}
}
