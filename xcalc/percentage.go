package xcalc

import (
	"fmt"
	"github.com/wbylovesun/xutils/types"
	"math"
	"strconv"
)

func numberPrecision[T types.GenericNumber](numerator, denominator T, precision int) (float64, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("denominator can not be 0")
	}
	n := float64(numerator)
	d := float64(denominator)
	precisionFormat := fmt.Sprintf("%%.%df", precision)
	return strconv.ParseFloat(fmt.Sprintf(precisionFormat, n/d), 64)
}

func NumberPrecision[T types.GenericNumber](numerator, denominator T, precision int) float64 {
	v, err := numberPrecision(numerator, denominator, precision)
	if err != nil {
		return 0
	}
	return v
}

// Percentage 默认输出带2位小数的比例，错误时报错
func Percentage[T types.GenericNumber](numerator, denominator T) (float64, error) {
	f, err := numberPrecision(numerator, denominator, 4)
	if err != nil {
		return 0, err
	}
	return PromiseWithTwoDigits(f), nil
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
	v, err := numberPrecision(numerator, denominator, precision+4)
	if err != nil {
		return 0, err
	}
	return Round(v*100, precision), nil
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

// PromisePercentWithTwoDigits 保留（最多）两位小数
// 参数是百分比的数值，譬如：25.64
func PromisePercentWithTwoDigits(n float64) float64 {
	return Round(n, 2)
}

// PromiseWithTwoDigits 保留（最多）两位小数
// 与PromisePercentWithTwoDigits不同点在于PromiseWithTwoDigits的参数是非百分比，譬如：
// PromisePercentWithTwoDigits(10)的参数是10%，而PromiseWithTwoDigits(10)的参数是10
// 如果这两者返回值是一样，则PromiseWithTwoDigits的参数应该是0.1
func PromiseWithTwoDigits(n float64) float64 {
	return Round(n*100, 2)
}

// PercentageWithTwoDigits 默认输出带2位小数的比例，忽略错误
// 根据给出的两个数字，计算出百分比，默认保留2位小数
// 使用的是PercentagePrecisionWithSuppression，本身会乘以100，因此需要使用PromisePercentWithTwoDigits
func PercentageWithTwoDigits[T types.GenericNumber](numerator, denominator T) float64 {
	v, err := Percentage(numerator, denominator)
	if err != nil {
		return 0
	}
	return v
}

// Round 带四舍五入指定精度的浮点数，ROUND_HALF_UP 模式实现
// 返回将 val 根据指定精度 precision（十进制小数点后数字的数目）进行四舍五入的结果。precision 也可以是负数或零。
// 参考：
// https://github.com/thinkeridea/go-extend/blob/main/exmath/round.go
func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

func Ratio(now, pre float64, precision int) float64 {
	if now == 0 && pre == 0 {
		return 0
	}
	if now == 0 {
		return -100
	}
	if pre == 0 {
		return math.Inf(0)
	}
	return Round((now/pre-1)*100, precision)
}

func PercentageRatio(now, pre float64) float64 {
	return Ratio(now, pre, 2)
}
