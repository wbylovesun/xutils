package xtime

import (
	"database/sql/driver"
	"errors"
	"time"
)

func (j *JsonShortDate) Value() (driver.Value, error) {
	if j.Time.IsZero() {
		return nil, nil
	}
	return j.Time, nil
}

func (j *JsonShortDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	v, ok := value.(time.Time)
	if ok {
		*j = JsonShortDate{v}
		return nil
	}
	return errors.New("can not convert %+v to timestamp")
}

func (j *JsonLongDate) Value() (driver.Value, error) {
	if j.Time.IsZero() {
		return nil, nil
	}
	return j.Time, nil
}

func (j *JsonLongDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	v, ok := value.(time.Time)
	if ok {
		*j = JsonLongDate{v}
		return nil
	}
	return errors.New("can not convert %+v to timestamp")
}
