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

type JsonShortDate struct {
	time.Time
}

type JsonLongDate struct {
	time.Time
}

type JsonTimestamp struct {
	time.Time
}
