package xvalidator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/wbylovesun/xutils/xtime"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var reTimeSpan = regexp.MustCompile(`([1-9]\d*)(year|y|month|M|day|d|hour|h|minute|m|second|s)`)

func timeSpan(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()
	params := strings.Split(param, ";")
	paramSize := len(params)
	if paramSize < 2 || paramSize > 3 {
		return false
	}
	fieldValue, ok := field.Interface().(time.Time)
	if !ok {
		return false
	}
	crossField := params[0]
	crossInf := fl.Parent().FieldByName(crossField).Interface()
	if crossInf == nil {
		return false
	}
	crossValue, ok := crossInf.(time.Time)
	if !ok {
		return false
	}

	for i := 1; i < paramSize; i++ {
		p := params[i]
		tsd := strings.Split(p, ":")
		if len(tsd) != 2 {
			return false
		}
		if !strings.Contains("gt,gte,lt,lte", tsd[0]) {
			return false
		}
		tst, err := parseTimeSpan(crossValue, tsd[1])
		if err != nil {
			return false
		}
		switch tsd[0] {
		case "gt":
			if !fieldValue.After(tst) {
				return false
			}
		case "lte":
			if fieldValue.After(tst) {
				return false
			}
		case "gte":
			if fieldValue.Before(tst) {
				return false
			}
		case "lt":
			if !fieldValue.Before(tst) {
				return false
			}
		}
	}

	return true
}

func parseTimeSpan(t time.Time, s string) (time.Time, error) {
	if !reTimeSpan.MatchString(s) {
		return time.Time{}, fmt.Errorf("invalid timespan format")
	}
	nt := t.Add(0)
	submatches := reTimeSpan.FindAllStringSubmatch(s, -1)
	for _, submatch := range submatches {
		atoi, err := strconv.Atoi(submatch[1])
		if err != nil {
			return time.Time{}, err
		}
		switch submatch[2] {
		case "y":
			fallthrough
		case "year":
			nt = nt.AddDate(atoi, 0, 0)
		case "M":
			fallthrough
		case "month":
			nt = nt.AddDate(0, atoi, 0)
		case "d":
			fallthrough
		case "day":
			nt = nt.AddDate(0, 0, atoi)
		case "h":
			fallthrough
		case "hour":
			nt = nt.Add(time.Duration(atoi) * time.Hour)
		case "m":
			fallthrough
		case "minute":
			nt = nt.Add(time.Duration(atoi) * time.Minute)
		case "s":
			fallthrough
		case "second":
			nt = nt.Add(time.Duration(atoi) * time.Second)
		default:
		}
	}
	return nt, nil
}

var reDateOnly = regexp.MustCompile(`^(today|yesterday|tomorrow|first day of this month|(-?\d+)\s*(year|month|week|day)s?)$`)
var reDateTime = regexp.MustCompile(`^(today|yesterday|tomorrow|(-?\d+)\s*(?:(year|month|week|day)s?)|first day of this month)\s*(\d{2}:\d{2}:\d{2})$`)

func tryParseDate(param string) (*time.Time, error) {
	expected, err := xtime.Parse(param)
	if err == nil {
		return &expected, nil
	}

	now := time.Now()

	if reDateOnly.MatchString(param) {
		submatches := reDateOnly.FindAllStringSubmatch(param, -1)
		submatch := submatches[0]
		if submatch[1] != "" {
			switch submatch[1] {
			case "today":
				expected = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
			case "yesterday":
				expected = now.AddDate(0, 0, -1)
			case "tomorrow":
				expected = now.AddDate(0, 0, 1)
			case "first day of this month":
				expected = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
			}
			return &expected, nil
		} else {
			gap, err := strconv.Atoi(submatch[2])
			if err != nil {
				return nil, err
			}
			switch submatch[3] {
			case "day":
				expected = now.AddDate(0, 0, gap)
				return &expected, nil
			case "week":
				expected = now.AddDate(0, 0, 7*gap)
				return &expected, nil
			case "month":
				expected = now.AddDate(0, gap, 0)
				return &expected, nil
			case "year":
				expected = now.AddDate(gap, 0, 0)
				return &expected, nil
			}
		}
	} else if reDateTime.MatchString(param) {
		submatches := reDateTime.FindAllStringSubmatch(param, -1)
		submatch := submatches[0]
		if submatch[2] != "" && submatch[3] != "" {
			gap, err := strconv.Atoi(submatch[1])
			if err != nil {
				return nil, err
			}
			switch submatch[3] {
			case "day":
				expected = now.AddDate(0, 0, gap)
			case "week":
				expected = now.AddDate(0, 0, 7*gap)
			case "month":
				expected = now.AddDate(0, gap, 0)
			case "year":
				expected = now.AddDate(gap, 0, 0)
			}
			ds := fmt.Sprintf("%s %s", expected.Format("2006-01-02"), submatch[4])
			d, err := time.Parse("2006-01-02 15:04:05", ds)
			return &d, err
		} else {
			switch submatch[1] {
			case "today":
				expected = time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
			case "yesterday":
				expected = now.AddDate(0, 0, -1)
			case "tomorrow":
				expected = now.AddDate(0, 0, 1)
			case "first day of this month":
				expected = time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
			}
			ds := fmt.Sprintf("%s %s", expected.Format("2006-01-02"), submatch[4])
			d, err := time.Parse("2006-01-02 15:04:05", ds)
			return &d, err
		}
	} else if duration, err := time.ParseDuration(param); err == nil {
		expected = now.Add(duration)
		return &expected, nil
	}
	return nil, fmt.Errorf("not supported format")
}

func parseParamToExpectedTime(param string) time.Time {
	now := time.Now()
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
	return expected
}
