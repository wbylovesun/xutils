package xtime

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"time"
)

func (j *DateTime) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte{'"'})
	buf.WriteString(j.String())
	buf.Write([]byte{'"'})
	return buf.Bytes(), nil
}

func (j *DateTime) UnmarshalJSON(data []byte) error {
	t, err := time.ParseInLocation(`"`+time.DateTime+`"`, string(data), time.Local)
	*j = DateTime(t)
	return err
}

func (j *DateTime) String(format ...string) string {
	layout := time.DateTime
	if len(format) > 0 {
		layout = format[0]
	}
	v := time.Time(*j)
	return v.Format(layout)
}

func (j *DateTime) Time() time.Time {
	v := time.Time(*j)
	return v
}

func (j *Date) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte{'"'})
	buf.WriteString(j.String())
	buf.Write([]byte{'"'})
	return buf.Bytes(), nil
}

func (j *Date) UnmarshalJSON(data []byte) error {
	t, err := time.ParseInLocation(`"`+time.DateOnly+`"`, string(data), time.Local)
	*j = Date(t)
	return err
}

func (j *Date) String(format ...string) string {
	layout := time.DateOnly
	if len(format) > 0 {
		layout = format[0]
	}
	v := time.Time(*j)
	return v.Format(layout)
}

func (j *Date) Time() time.Time {
	v := time.Time(*j)
	return v
}

func (j *JsonTimestamp) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	v := time.Time(*j)
	binary.Write(buf, binary.LittleEndian, v.Unix())
	return buf.Bytes(), nil
}

func (j *JsonTimestamp) UnmarshalJSON(data []byte) error {
	atoi, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	t := time.Unix(int64(atoi), 0)
	*j = JsonTimestamp(t)
	return nil
}

func (j *JsonTimestamp) String(format ...string) string {
	layout := time.DateTime
	if len(format) > 0 {
		layout = format[0]
	}
	v := time.Time(*j)
	return v.Format(layout)
}

func (j *JsonTimestamp) Time() time.Time {
	v := time.Time(*j)
	return v
}
