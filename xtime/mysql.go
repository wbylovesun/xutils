package xtime

import (
	"database/sql/driver"
	"errors"
	"time"
)

func (j Date) Value() (driver.Value, error) {
	if j.Time().IsZero() {
		return nil, nil
	}
	return j.Time(), nil
}

func (j *Date) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	v, ok := value.(time.Time)
	if ok {
		*j = Date(v)
		return nil
	}
	return errors.New("can not convert %+v to timestamp")
}

func (j DateTime) Value() (driver.Value, error) {
	if j.Time().IsZero() {
		return nil, nil
	}
	return j.Time(), nil
}

func (j *DateTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	v, ok := value.(time.Time)
	if ok {
		*j = DateTime(v)
		return nil
	}
	return errors.New("can not convert %+v to timestamp")
}

func (j JsonTimestamp) Value() (driver.Value, error) {
	if j.Time().IsZero() {
		return nil, nil
	}
	return j.Time().Unix(), nil
}

func (j *JsonTimestamp) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	v, ok := value.(time.Time)
	if ok {
		*j = JsonTimestamp(v)
		return nil
	}
	return errors.New("can not convert %+v to timestamp")
}
