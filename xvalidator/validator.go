package xvalidator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"sync"
	"time"
)

var (
	timeDurationType = reflect.TypeOf(time.Duration(0))
	timeType         = reflect.TypeOf(time.Time{})
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

//var _ binding.StructValidator = &defaultValidator{}

//var timeMark = []string{"today", "week", "month", "yesterday", "last_week", "last_month"}

func (v *defaultValidator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
		_ = v.validate.RegisterValidation("int_slice", intSlice)
		_ = v.validate.RegisterValidation("gte", isGte)
		_ = v.validate.RegisterValidation("gt", isGt)
		_ = v.validate.RegisterValidation("lte", isLte)
		_ = v.validate.RegisterValidation("lt", isLt)
		_ = v.validate.RegisterValidation("time_span", timeSpan)

	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}

func NewValidator() *defaultValidator {
	return new(defaultValidator)
}

// asInt returns the parameter as a int64
// or panics if it can't convert
func asInt(param string) int64 {
	i, err := strconv.ParseInt(param, 0, 64)
	panicIf(err)

	return i
}

// asIntFromTimeDuration parses param as time.Duration and returns it as int64
// or panics on error.
func asIntFromTimeDuration(param string) int64 {
	d, err := time.ParseDuration(param)
	if err != nil {
		// attempt parsing as an an integer assuming nanosecond precision
		return asInt(param)
	}
	return int64(d)
}

// asIntFromType calls the proper function to parse param as int64,
// given a field's Type t.
func asIntFromType(t reflect.Type, param string) int64 {
	switch t {
	case timeDurationType:
		return asIntFromTimeDuration(param)
	default:
		return asInt(param)
	}
}

// asUint returns the parameter as a uint64
// or panics if it can't convert
func asUint(param string) uint64 {

	i, err := strconv.ParseUint(param, 0, 64)
	panicIf(err)

	return i
}

// asFloat returns the parameter as a float64
// or panics if it can't convert
func asFloat(param string) float64 {

	i, err := strconv.ParseFloat(param, 64)
	panicIf(err)

	return i
}

// asBool returns the parameter as a bool
// or panics if it can't convert
func asBool(param string) bool {

	i, err := strconv.ParseBool(param)
	panicIf(err)

	return i
}

func panicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}
