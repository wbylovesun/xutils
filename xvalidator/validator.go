package xvalidator

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"sync"
	"time"
)

var (
	timeDurationType = reflect.TypeOf(time.Duration(0))
	timeType         = reflect.TypeOf(time.Time{})
)

type Validator struct {
	once     sync.Once
	validate *validator.Validate
}

//var _ binding.StructValidator = &Validator{}

//var timeMark = []string{"today", "week", "month", "yesterday", "last_week", "last_month"}

func (v *Validator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyInit()

		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}

	return nil
}

func (v *Validator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

func (v *Validator) lazyInit() {
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

func NewValidator() *Validator {
	return new(Validator)
}
