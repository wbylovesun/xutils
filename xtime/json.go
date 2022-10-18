package xtime

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"time"
)

func (j *JsonLongDate) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte{'"'})
	buf.WriteString(j.String())
	buf.Write([]byte{'"'})
	return buf.Bytes(), nil
}

func (j *JsonLongDate) UnmarshalJSON(data []byte) error {
	t, err := time.ParseInLocation(`"`+LongFormat+`"`, string(data), time.Local)
	*j = JsonLongDate{t}
	return err
}

func (j *JsonLongDate) String(format ...string) string {
	layout := LongFormat
	if len(format) > 0 {
		layout = format[0]
	}
	return j.Format(layout)
}

func (j *JsonShortDate) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	buf.Write([]byte{'"'})
	buf.WriteString(j.String())
	buf.Write([]byte{'"'})
	return buf.Bytes(), nil
}

func (j *JsonShortDate) UnmarshalJSON(data []byte) error {
	t, err := time.ParseInLocation(`"`+ShortFormat+`"`, string(data), time.Local)
	*j = JsonShortDate{t}
	return err
}

func (j *JsonShortDate) String(format ...string) string {
	layout := ShortFormat
	if len(format) > 0 {
		layout = format[0]
	}
	return j.Format(layout)
}

func (j *JsonTimestamp) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, j.Unix())
	return buf.Bytes(), nil
}

func (j *JsonTimestamp) UnmarshalJSON(data []byte) error {
	atoi, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	t := time.Unix(int64(atoi), 0)
	*j = JsonTimestamp{t}
	return nil
}

func (j *JsonTimestamp) String(format ...string) string {
	layout := LongFormat
	if len(format) > 0 {
		layout = format[0]
	}
	return j.Format(layout)
}
