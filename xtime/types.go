package xtime

import "time"

type DateRange struct {
	from time.Time
	to   time.Time
}

type ComparableDateRange struct {
	DateRange
}

type WeekRange struct {
	from time.Time
	to   time.Time
}

type ComparableWeekRange struct {
	WeekRange
}

type MonthRange struct {
	from time.Time
	to   time.Time
}

type ComparableMonthRange struct {
	MonthRange
}

type JsonShortDate time.Time

type JsonLongDate time.Time

type JsonTimestamp time.Time
