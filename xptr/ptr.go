package xptr

// Ptr 返回指向 val 副本的指针。
func Ptr[V any](val V) *V {
	return &val
}

type canConvertIntPointerTypes interface {
	int8 | int16 | int32 | int64 | int | float64 | float32
}

// PtrToIntPtr 将数值类型指针转换为 *int，nil 透传。
// 转换遵循 Go 语义：整数溢出回绕，float 小数部分截断。
func PtrToIntPtr[T canConvertIntPointerTypes](dt *T) *int {
	if dt == nil {
		return nil
	}
	dtInt := int(*dt)
	return &dtInt
}

// PtrValueOrEmpty 解引用指针，nil 时返回 V 的零值。
func PtrValueOrEmpty[V any, T *V](val T) V {
	if val == nil {
		return *new(V)
	}

	return *val
}

// PtrValueOrDefault 解引用指针，nil 时返回默认值 def。
func PtrValueOrDefault[V any, T *V](val T, def V) V {
	if val == nil {
		return def
	}

	return *val
}

// IntPtrTo 将 *int 转换为 *T，nil 透传。
// 数值转换遵循 Go 语义：整数溢出回绕，float 小数部分截断。
func IntPtrTo[T canConvertIntPointerTypes](dt *int) *T {
	return PtrTo[int, T](dt)
}

// PtrTo 将数值类型指针 *T 转换为 *O，nil 透传。
// 数值转换遵循 Go 语义：整数溢出回绕，float 小数部分截断。
func PtrTo[T, O canConvertIntPointerTypes](dt *T) *O {
	if dt == nil {
		return nil
	}
	o := O(*dt)
	return &o
}

// ClonePtr 返回指向值副本的新指针，nil 透传。
func ClonePtr[V any](val *V) *V {
	if val == nil {
		return nil
	}
	v := *val
	return &v
}

// EqualPtr 比较两个指针指向的值：均为 nil 视为相等，仅一者为 nil 不相等，否则比较指向的值。
func EqualPtr[V comparable](a, b *V) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}

// Coalesce 返回第一个非 nil 的指针（返回原指针，不做复制），全部为 nil 时返回 nil。
func Coalesce[T any](vals ...*T) *T {
	for _, v := range vals {
		if v != nil {
			return v
		}
	}
	return nil
}

// CoalesceValue 返回第一个非 nil 指针指向的值，全部为 nil 时返回 def。
func CoalesceValue[T any](def T, vals ...*T) T {
	for _, v := range vals {
		if v != nil {
			return *v
		}
	}
	return def
}
