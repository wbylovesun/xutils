package xvalidator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"time"
	"unicode/utf8"
)

// isGte is the validation function for validating if the current field's value is greater than or equal to the param's value.
func isGte(fl validator.FieldLevel) bool {

	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) >= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) >= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asIntFromType(field.Type(), param)

		return field.Int() >= p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() >= p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() >= p

	case reflect.Struct:

		if field.Type() == timeType {
			now := time.Now().UTC()
			var expected time.Time
			switch param {
			//case "today":
			//	expected = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
			//case "yesterday":
			//	expected = now.AddDate(0, 0, -1)
			//case "tomorrow":
			//	expected = now.AddDate(0, 0, 1)
			case "":
				expected = now
			default:
				d, err := tryParseDate(param)
				if err != nil {
					expected = now
				} else {
					expected = *d
				}
			}

			t := field.Interface().(time.Time)

			return t.After(expected) || t.Equal(expected)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// isGt is the validation function for validating if the current field's value is greater than the param's value.
func isGt(fl validator.FieldLevel) bool {

	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) > p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) > p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asIntFromType(field.Type(), param)

		return field.Int() > p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() > p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() > p
	case reflect.Struct:
		if field.Type() == timeType {
			now := time.Now().UTC()
			var expected time.Time
			switch param {
			case "":
				expected = now
			default:
				d, err := tryParseDate(param)
				if err != nil {
					expected = now
				} else {
					expected = *d
				}
			}

			t := field.Interface().(time.Time)

			return t.After(expected)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// isLte is the validation function for validating if the current field's value is less than or equal to the param's value.
func isLte(fl validator.FieldLevel) bool {

	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) <= p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) <= p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asIntFromType(field.Type(), param)

		return field.Int() <= p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() <= p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() <= p

	case reflect.Struct:

		if field.Type() == timeType {

			now := time.Now().UTC()
			var expected time.Time
			switch param {
			case "":
				expected = now
			default:
				d, err := tryParseDate(param)
				if err != nil {
					expected = now
				} else {
					expected = *d
				}
			}

			t := field.Interface().(time.Time)

			return t.Before(expected) || t.Equal(expected)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

// isLt is the validation function for validating if the current field's value is less than the param's value.
func isLt(fl validator.FieldLevel) bool {

	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {

	case reflect.String:
		p := asInt(param)

		return int64(utf8.RuneCountInString(field.String())) < p

	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)

		return int64(field.Len()) < p

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p := asIntFromType(field.Type(), param)

		return field.Int() < p

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p := asUint(param)

		return field.Uint() < p

	case reflect.Float32, reflect.Float64:
		p := asFloat(param)

		return field.Float() < p

	case reflect.Struct:

		if field.Type() == timeType {
			now := time.Now().UTC()
			var expected time.Time
			switch param {
			case "":
				expected = now
			default:
				d, err := tryParseDate(param)
				if err != nil {
					expected = now
				} else {
					expected = *d
				}
			}

			t := field.Interface().(time.Time)
			return t.Before(expected)
		}
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}
