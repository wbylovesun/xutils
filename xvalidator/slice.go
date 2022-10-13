package xvalidator

import (
	"github.com/go-playground/validator/v10"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var reIntSlice = regexp.MustCompile(`(lt|gt|lte|gte|sep):(\w+);?`)

// 验证以separator为分隔符的正整数集合
// 当分隔的部分使用strconv.Atoi转换失败时，验证不通过
// Example:
// binding:"intSlice" 默认使用","作为分隔符
// binding:"intSlice=sep:0x2C;gt:0;lte:100"
func intSlice(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	param := fl.Param()

	var (
		separator                    = ","
		min                          = math.MinInt64
		max                          = math.MaxInt64
		hasLt, hasLte, hasGt, hasGte = false, false, false, false
		err                          error
	)
	if param != "" {
		if reIntSlice.MatchString(param) {
			submatches := reIntSlice.FindAllStringSubmatch(param, -1)
			for _, submatch := range submatches {
				err = nil
				if submatch[1] == "sep" {
					separator = submatch[2]
				} else if submatch[1] == "lt" {
					hasLt = true
					max, err = strconv.Atoi(submatch[2])
				} else if submatch[1] == "lte" {
					hasLte = true
					max, err = strconv.Atoi(submatch[2])
				} else if submatch[1] == "gt" {
					hasGt = true
					min, err = strconv.Atoi(submatch[2])
				} else if submatch[1] == "gte" {
					hasGte = true
					min, err = strconv.Atoi(submatch[2])
				}
				if err != nil {
					return false
				}
			}
			if hasLt && hasLte {
				return false
			}
			if hasGt && hasGte {
				return false
			}
			if min > max {
				return false
			}
		} else {
			separator = param[0:1]
		}
	}
	splits := strings.Split(fl.Field().String(), separator)
	for _, str := range splits {
		v, err := strconv.Atoi(str)
		if err != nil {
			return false
		}
		if hasLt && v >= max {
			return false
		} else if hasLte && v > max {
			return false
		} else if hasGt && v <= min {
			return false
		} else if hasGte && v < min {
			return false
		}
	}
	return true
}
