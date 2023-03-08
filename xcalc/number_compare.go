package xcalc

import (
	"github.com/wbylovesun/xutils/types"
	"math"
)

// Min Return minimum value in z.
// T(0) returned if non-exists in z.
// Value of index(0) if only 1 item in z.
func Min[T types.GenericNumber](z ...T) T {
	if len(z) == 0 {
		return T(0)
	}
	if len(z) == 1 {
		return z[0]
	}
	val := T(math.Inf(1))
	for i := 0; i < len(z); i++ {
		if val > z[i] {
			val = z[i]
		}
	}
	return val
}

// Max Return maximum value in z.
// T(0) returned if non-exists in z.
// Value of index(0) if only 1 item in z.
func Max[T types.GenericNumber](z ...T) T {
	if len(z) == 0 {
		return T(0)
	}
	if len(z) == 1 {
		return z[0]
	}
	val := T(math.Inf(-1))
	for i := 0; i < len(z); i++ {
		if z[i] > val {
			val = z[i]
		}
	}
	return val
}

// Gte0
//
// Deprecated: Should use PromiseGte0 instead
func Gte0[T types.GenericNumber](x T) T {
	return Gte(x, 0)
}

// Gte
//
// Deprecated: Should use PromiseGte instead
func Gte[T types.GenericNumber](x, y T) T {
	if x < y {
		return y
	}
	return x
}

// PromiseGte0 Promise value of x must be greater than or equal 0
func PromiseGte0[T types.GenericNumber](x T) T {
	return PromiseGte(x, 0)
}

// PromiseGte Promise value of x must be greater than or equal y
func PromiseGte[T types.GenericNumber](x, y T) T {
	return Max(x, y)
}

// Lte100
//
// Deprecated: Should use PromiseLte100 instead
func Lte100[T types.GenericNumber](x T) T {
	return Lte(x, 100)
}

// Lte
//
// Deprecated: Should use PromiseLte instead
func Lte[T types.GenericNumber](x, y T) T {
	if x > y {
		return y
	}
	return x
}

// PromiseLte100 Promise value of x must be less than or equal 100
func PromiseLte100[T types.GenericNumber](x T) T {
	return PromiseLte(x, 100)
}

// PromiseLte Promise value of x must be less than or equal y
func PromiseLte[T types.GenericNumber](x, y T) T {
	return Min(x, y)
}

// IsGte Check x is greater than or equal y or not
func IsGte[T types.GenericNumber](x, y T) bool {
	return x >= y
}

// IsLte Check x is less than or equal y or not
func IsLte[T types.GenericNumber](x, y T) bool {
	return x <= y
}

// IsGt Check x is greater than y or not
func IsGt[T types.GenericNumber](x, y T) bool {
	return x > y
}

// IsLt Check x is less than y or not
func IsLt[T types.GenericNumber](x, y T) bool {
	return x < y
}
