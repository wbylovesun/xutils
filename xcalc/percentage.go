package xcalc

import (
	"fmt"
	"github.com/wbylovesun/xutils/types"
	"strconv"
)

func Percentage[T types.GenericNumber](numerator, denominator T) (float64, error) {
	return PercentagePrecision(numerator, denominator, 2)
}

func PercentageNoDigit[T types.GenericNumber](numerator, denominator T) (int, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("dominator can not be 0")
	}
	n := float64(numerator)
	d := float64(denominator)
	return int(n / d * 100), nil
}

func PercentagePrecision[T types.GenericNumber](numerator, denominator T, precision int) (float64, error) {
	if denominator == 0 {
		return 0, fmt.Errorf("denominator can not be 0")
	}
	n := float64(numerator)
	d := float64(denominator)
	precisionFormat := fmt.Sprintf("%%.%df", precision)
	return strconv.ParseFloat(fmt.Sprintf(precisionFormat, n/d*100), 64)
}

func PercentageWithSuppression[T types.GenericNumber](numerator, denominator T) float64 {
	f, _ := PercentagePrecision(numerator, denominator, 0)
	return f
}

func PercentageNoDigitWithSuppression[T types.GenericNumber](numerator, denominator T) int {
	f, _ := PercentageNoDigit(numerator, denominator)
	return f
}

func PercentagePrecisionWithSuppression[T types.GenericNumber](numereator, denominator T, precision int) float64 {
	f, _ := PercentagePrecision(numereator, denominator, precision)
	return f
}
