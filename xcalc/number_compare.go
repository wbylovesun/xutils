package xcalc

import "github.com/wbylovesun/xutils/types"

func Min[T types.GenericNumber](x, y T) T {
	if x > y {
		return y
	}
	return x
}

func Max[T types.GenericNumber](x, y T) T {
	if x > y {
		return x
	}
	return y
}

func Gte0[T types.GenericNumber](x T) T {
	return Gte(x, 0)
}

func Gte[T types.GenericNumber](x, y T) T {
	if x < y {
		return y
	}
	return x
}

func Lte100[T types.GenericNumber](x T) T {
	return Lte(x, 100)
}

func Lte[T types.GenericNumber](x, y T) T {
	if x > y {
		return y
	}
	return x
}

func IsGte[T types.GenericNumber](x, y T) bool {
	return x >= y
}

func IsLte[T types.GenericNumber](x, y T) bool {
	return x <= y
}

func IsGt[T types.GenericNumber](x, y T) bool {
	return x > y
}

func IsLt[T types.GenericNumber](x, y T) bool {
	return x < y
}
