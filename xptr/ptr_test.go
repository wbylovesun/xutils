package xptr

import (
	"reflect"
	"testing"
)

func TestPtr(t *testing.T) {
	p := Ptr(42)
	if p == nil {
		t.Fatal("Ptr(42) 返回了 nil")
	}
	if *p != 42 {
		t.Errorf("Ptr(42) 解引用 = %d, want 42", *p)
	}

	sp := Ptr("hello")
	if sp == nil || *sp != "hello" {
		t.Errorf(`Ptr("hello") = %v, want 指向 "hello" 的指针`, sp)
	}

	// 每次调用应返回独立变量的地址
	a, b := Ptr(1), Ptr(1)
	if a == b {
		t.Error("两次 Ptr(1) 返回了同一指针，应为相互独立的变量")
	}
	if *a != 1 || *b != 1 {
		t.Errorf("两次 Ptr(1) 解引用 = %d, %d, want 1, 1", *a, *b)
	}
}

func TestPtrToIntPtr(t *testing.T) {
	t.Run("nil 输入返回 nil", func(t *testing.T) {
		if got := PtrToIntPtr[int](nil); got != nil {
			t.Errorf("PtrToIntPtr[int](nil) = %v, want nil", got)
		}
	})

	t.Run("int", func(t *testing.T) {
		v := 42
		got := PtrToIntPtr(&v)
		if got == nil || *got != 42 {
			t.Errorf("PtrToIntPtr(&42) = %v, want 42", got)
		}
	})

	t.Run("int8 边界", func(t *testing.T) {
		v := int8(-128)
		got := PtrToIntPtr(&v)
		if got == nil || *got != -128 {
			t.Errorf("PtrToIntPtr(&int8(-128)) = %v, want -128", got)
		}
	})

	t.Run("int32 大值", func(t *testing.T) {
		v := int32(2147483647)
		got := PtrToIntPtr(&v)
		if got == nil || *got != 2147483647 {
			t.Errorf("PtrToIntPtr(&int32 max) = %v, want 2147483647", got)
		}
	})

	t.Run("int64", func(t *testing.T) {
		v := int64(1) << 40
		got := PtrToIntPtr(&v)
		if got == nil || *got != 1<<40 {
			t.Errorf("PtrToIntPtr(&int64) = %v, want %d", got, int64(1)<<40)
		}
	})

	t.Run("float 整数值", func(t *testing.T) {
		f64 := 3.0
		if got := PtrToIntPtr(&f64); got == nil || *got != 3 {
			t.Errorf("PtrToIntPtr(&3.0) = %v, want 3", got)
		}
		f32 := float32(-2.0)
		if got := PtrToIntPtr(&f32); got == nil || *got != -2 {
			t.Errorf("PtrToIntPtr(&float32(-2.0)) = %v, want -2", got)
		}
	})
}

func TestPtrValueOrEmpty(t *testing.T) {
	v := 5
	if got := PtrValueOrEmpty[int](&v); got != 5 {
		t.Errorf("PtrValueOrEmpty(&5) = %d, want 5", got)
	}
	if got := PtrValueOrEmpty[int, *int](nil); got != 0 {
		t.Errorf("PtrValueOrEmpty(nil) = %d, want 零值 0", got)
	}

	s := "abc"
	if got := PtrValueOrEmpty[string](&s); got != "abc" {
		t.Errorf("PtrValueOrEmpty(&%q) = %q, want %q", s, got, "abc")
	}
	if got := PtrValueOrEmpty[string, *string](nil); got != "" {
		t.Errorf("PtrValueOrEmpty[string](nil) = %q, want 空字符串", got)
	}

	type st struct{ A int }
	if got := PtrValueOrEmpty[st, *st](nil); !reflect.DeepEqual(got, st{}) {
		t.Errorf("PtrValueOrEmpty[st](nil) = %+v, want 零值结构体", got)
	}
}

func TestPtrValueOrDefault(t *testing.T) {
	v := 5
	if got := PtrValueOrEmpty[int](&v); got != 5 {
		t.Errorf("PtrValueOrDefault(&5) = %d, want 5", got)
	}
	if got := PtrValueOrDefault[int, *int](nil, 7); got != 7 {
		t.Errorf("PtrValueOrDefault(nil, 7) = %d, want 7", got)
	}
	if got := PtrValueOrDefault[int](&v, 7); got != 5 {
		t.Errorf("PtrValueOrDefault(&5, 7) = %d, want 5（非 nil 时应忽略默认值）", got)
	}

	if got := PtrValueOrDefault[string, *string](nil, "def"); got != "def" {
		t.Errorf("PtrValueOrDefault[string](nil, %q) = %q, want %q", "def", got, "def")
	}
}

func TestIntPtrTo(t *testing.T) {
	if got := IntPtrTo[int8](nil); got != nil {
		t.Errorf("IntPtrTo[int8](nil) = %v, want nil", got)
	}

	v := 42
	if got := IntPtrTo[int8](&v); got == nil || *got != 42 {
		t.Errorf("IntPtrTo[int8](&42) = %v, want 42", got)
	}

	neg := -42
	if got := IntPtrTo[int16](&neg); got == nil || *got != -42 {
		t.Errorf("IntPtrTo[int16](&-42) = %v, want -42", got)
	}
}

func TestPtrTo(t *testing.T) {
	if got := PtrTo[int, int8](nil); got != nil {
		t.Errorf("PtrTo[int, int8](nil) = %v, want nil", got)
	}

	i := 7
	if got := PtrTo[int, int8](&i); got == nil || *got != 7 {
		t.Errorf("PtrTo[int, int8](&7) = %v, want 7", got)
	}

	// float → int 遵循 Go 语义：截断小数部分
	f := 3.5
	if got := PtrTo[float64, int](&f); got == nil || *got != 3 {
		t.Errorf("PtrTo[float64, int](&3.5) = %v, want 3", got)
	}

	// int → float
	if got := PtrTo[int, float64](&i); got == nil || *got != 7 {
		t.Errorf("PtrTo[int, float64](&7) = %v, want 7", got)
	}
}

func TestClonePtr(t *testing.T) {
	if got := ClonePtr[int](nil); got != nil {
		t.Errorf("ClonePtr[int](nil) = %v, want nil", got)
	}

	v := 5
	got := ClonePtr(&v)
	if got == nil || *got != 5 {
		t.Fatalf("ClonePtr(&5) = %v, want 指向 5", got)
	}
	if got == &v {
		t.Error("ClonePtr 返回了原指针，应指向副本")
	}

	*got = 100
	if v != 5 {
		t.Errorf("修改 ClonePtr 结果影响了原值：v = %d, want 5", v)
	}
}

func TestEqualPtr(t *testing.T) {
	a, b := 1, 1
	if !EqualPtr(&a, &b) {
		t.Error("EqualPtr(&1, &1) = false, want true（不同地址相同值）")
	}

	c := 2
	if EqualPtr(&a, &c) {
		t.Error("EqualPtr(&1, &2) = true, want false")
	}

	if !EqualPtr[int](nil, nil) {
		t.Error("EqualPtr(nil, nil) = false, want true")
	}

	if EqualPtr[int](nil, &a) || EqualPtr[int](&a, nil) {
		t.Error("仅一者为 nil 时应返回 false")
	}
}

func TestCoalesce(t *testing.T) {
	a, b := 1, 2
	if got := Coalesce[int](nil, &a, &b); got != &a {
		t.Errorf("Coalesce(nil, &a, &b) = %v, want 原指针 &a", got)
	}
	if got := Coalesce[int](nil, nil); got != nil {
		t.Errorf("Coalesce(nil, nil) = %v, want nil", got)
	}
	if got := Coalesce[int](); got != nil {
		t.Errorf("Coalesce() = %v, want nil", got)
	}
}

func TestCoalesceValue(t *testing.T) {
	a, b := 1, 2
	if got := CoalesceValue(0, nil, &a, &b); got != 1 {
		t.Errorf("CoalesceValue(0, nil, &a, &b) = %d, want 1", got)
	}
	if got := CoalesceValue(9, nil, nil); got != 9 {
		t.Errorf("CoalesceValue(9, nil, nil) = %d, want 9", got)
	}
	if got := CoalesceValue(9); got != 9 {
		t.Errorf("CoalesceValue(9) = %d, want 9", got)
	}
}
