package xcalc

import (
	"fmt"
	"github.com/wbylovesun/xutils/types"
	"math"
	"strconv"
)

// Percentage 默认输出带2位小数的比例，错误时报错
func Percentage[T types.GenericNumber](numerator, denominator T) (float64, error) {
	return PercentagePrecision(numerator, denominator, 2)
}

// PercentageNoDigit 不带小数，错误时报错
func PercentageNoDigit[T types.GenericNumber](numerator, denominator T) (int, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("dominator can not be 0")
	}
	n := float64(numerator)
	d := float64(denominator)
	return int(n / d * 100), nil
}

// PercentagePrecision 自定义小数位长度，错误时报错
func PercentagePrecision[T types.GenericNumber](numerator, denominator T, precision int) (float64, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("denominator can not be 0")
	}
	n := float64(numerator)
	d := float64(denominator)
	precisionFormat := fmt.Sprintf("%%.%df", precision)
	return strconv.ParseFloat(fmt.Sprintf(precisionFormat, n/d*100), 64)
}

// PercentageWithSuppression 带2位小数，抑制错误
func PercentageWithSuppression[T types.GenericNumber](numerator, denominator T) float64 {
	return PercentagePrecisionWithSuppression(numerator, denominator, 2)
}

// PercentageNoDigitWithSuppression 不带小数，抑制错误
func PercentageNoDigitWithSuppression[T types.GenericNumber](numerator, denominator T) int {
	if math.IsNaN(float64(numerator)) || math.IsNaN(float64(denominator)) {
		return 0
	}
	if math.IsInf(float64(numerator), 0) || math.IsInf(float64(denominator), 0) {
		return 0
	}
	f, err := PercentageNoDigit(numerator, denominator)
	if err != nil || math.IsNaN(float64(f)) || math.IsInf(float64(f), 0) {
		return 0
	}
	return f
}

// PercentagePrecisionWithSuppression 自定义小数位长度，抑制错误
func PercentagePrecisionWithSuppression[T types.GenericNumber](numerator, denominator T, precision int) float64 {
	if math.IsNaN(float64(numerator)) || math.IsNaN(float64(denominator)) {
		return 0
	}
	if math.IsInf(float64(numerator), 0) || math.IsInf(float64(denominator), 0) {
		return 0
	}
	f, err := PercentagePrecision(numerator, denominator, precision)
	if err != nil || math.IsNaN(f) || math.IsInf(f, 0) {
		return float64(0)
	}
	return f
}
