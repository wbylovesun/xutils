package xmath

import "math"

func Round(x float64) int {
	if math.IsNaN(x) {
		return 0
	}
	if math.IsInf(x, 0) {
		return 0
	}
	return int(math.Round(x))
}
