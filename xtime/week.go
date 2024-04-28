package xtime

import (
	"fmt"
	"time"
)

// LastWeekSameDay 返回上周同一天的时间
func LastWeekSameDay(t time.Time) time.Time {
	return AddDays(t, -7)
}

// WeekDateRange 返回一周的日期范围，且时间均为00:00:00
// 与ISOWeekDateRange不同，此函数的每周第一天为星期日
func WeekDateRange(t time.Time) DateRange {
	s := WithTime(AddDays(t, -1*int(t.Weekday())), 0, 0, 0)
	e := WithTime(AddDays(s, 6), 0, 0, 0)
	fmt.Println(s, e)
	return DateRange{
		from: s,
		to:   e,
	}
}

// StartOfWeek 返回一周的开始日期，且时间均为00:00:00
// 与ISOStartOfWeek不同，此函数每周的第一天为星期日
func StartOfWeek(t time.Time) time.Time {
	return WithTime(AddDays(t, -1*int(t.Weekday())), 0, 0, 0)
}

// EndOfWeek 返回一周的最后一天，且时间均为00:00:00
// 与ISOEndOfWeek不同，此函数的每周第一天为星期天
func EndOfWeek(t time.Time) time.Time {
	return WithTime(AddDays(t, 6-(int(t.Weekday()))), 0, 0, 0)
}
