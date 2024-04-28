package xtime

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func ISOInfo(t time.Time) string {
	year, week := t.ISOWeek()
	weekday := ISOWeekday(t)
	return fmt.Sprintf("%4d-W%02d-%d", year, week, weekday)
}

func TimeOfISOInfo(info string) (time.Time, error) {
	year, week, weekday, err := parseISOInfo(info)
	if err != nil {
		return time.Time{}, err
	}
	return ISOStartOfYear(MustOf(fmt.Sprintf("%d-01-04", year))).AddDate(0, 0, (week-1)*7+weekday-1), nil
}

func ISOStartOfYear(t time.Time) time.Time {
	year, _, _ := t.Date()
	return ISOStartOfWeek(MustOf(fmt.Sprintf("%d-01-04", year)))
}

func ISOEndOfYear(year int) time.Time {
	weeks := ISOWeeksInYear(year)
	t, _ := TimeOfISOInfo(fmt.Sprintf("%d-W%02d-7", year, weeks))
	return t
}

var pattern = regexp.MustCompile(`^(\d{4})-W(\d{1,2})-(\d)$`)

func parseISOInfo(info string) (year, week, weekday int, err error) {
	matches := pattern.FindStringSubmatch(info)
	fmt.Println(matches)
	if len(matches) != 4 {
		err = errors.New("invalid iso info")
		return -1, -1, -1, err
	}
	year, err = strconv.Atoi(matches[1])
	week, err = strconv.Atoi(matches[2])
	weekday, err = strconv.Atoi(matches[3])
	if year < 2000 || year > 2099 {
		return -1, -1, -1, errors.New("invalid year")
	}
	if week < 1 || week > ISOWeeksInYear(year) {
		return -1, -1, -1, errors.New("invalid week")
	}
	if weekday < 1 || weekday > 7 {
		return -1, -1, -1, errors.New("invalid weekday")
	}
	return year, week, weekday, nil
}

// ISOYearWeek 当前时间所在年份及周数
// ISO格式为2022-W52-1，本函数返回的是202252
func ISOYearWeek(t time.Time) int {
	year, week := t.ISOWeek()
	return year*100 + week
}

// ISOWeekday 当前时间所在周几
func ISOWeekday(t time.Time) int {
	weekday := t.Weekday()
	if weekday == 0 {
		weekday = 7
	}
	return int(weekday)
}

// TimeOfISOWeekday 返回指定周几的日期，时间不变，只对日期进行加减处理
func TimeOfISOWeekday(t time.Time, weekday int) time.Time {
	if weekday <= 0 || weekday > 7 {
		return t
	}
	day := ISOWeekday(t)
	if weekday == day {
		return t
	}
	gap := day - weekday
	if gap < 0 {
		gap += 8
	} else {
		gap = -1 * gap
	}
	return t.AddDate(0, 0, gap)
}

// ISOWeeksInYear 返回一年中的周数
func ISOWeeksInYear(year int) int {
	t := time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC)
	if t.Weekday() == time.Thursday {
		return 53
	}
	if IsLeap(year) && t.Weekday() == time.Friday {
		return 53
	}
	return 52
}

// ISOWeekDateRange 返回一周的日期范围，每周从周一开始，且时间均为00:00:00
func ISOWeekDateRange(t time.Time) DateRange {
	s := WithTime(AddDays(t, -1*((int(t.Weekday())+6)%7)), 0, 0, 0)
	e := WithTime(AddDays(s, 6), 0, 0, 0)
	return DateRange{
		from: s,
		to:   e,
	}
}

// ISOStartOfWeek 返回一周的第一天，每周从周一开始，且时间均为00:00:00
func ISOStartOfWeek(t time.Time) time.Time {
	return WithTime(AddDays(t, -1*((int(t.Weekday())+6)%7)), 0, 0, 0)
}

// ISOEndOfWeek 返回一周的最后一天，每周从周一开始，且时间均为00:00:00
func ISOEndOfWeek(t time.Time) time.Time {
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return WithTime(AddDays(t, 7-weekday), 0, 0, 0)
}
