package xcalc

import (
	"fmt"
	"github.com/wbylovesun/xutils/types"
	"reflect"
	"strconv"
)

type AmountType interface {
	types.Number
}

// FormatAmountToYuan 格式化金额为元，保留2位小数
//
// 当数值不为float32/float64类型时，单位会视为分，除以100后得到元
//
// 当数值为float32/float64时，仅格式化小数点后2位
func FormatAmountToYuan[T AmountType](amount T) float64 {
	var yuan float64
	centKind := reflect.ValueOf(amount).Kind()
	if centKind != reflect.Float32 && centKind != reflect.Float64 {
		yuan = float64(amount) / 100
	} else if centKind == reflect.Float32 {
		yuan = float64(amount)
	} else {
		yuan = float64(amount)
	}
	formatted, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", yuan), 64)
	return formatted
}
