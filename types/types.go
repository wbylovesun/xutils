package types

// SignedNumber 有符号整型，含底层类型为有符号整型的自定义类型。
type SignedNumber interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedNumber 无符号整型，含底层类型为无符号整型的自定义类型。
type UnsignedNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Integer 全部整型（有符号 + 无符号）。
type Integer interface {
	SignedNumber | UnsignedNumber
}

// Number 全量数值类型（整型 + 浮点型），含底层类型为数值的自定义类型。
// 与标准库 cmp.Ordered 的分工：需要比较大小或包含 string 时用 cmp.Ordered，纯数值算术/格式化用 Number。
type Number interface {
	SignedNumber | UnsignedNumber | ~float32 | ~float64
}
