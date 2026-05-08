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

type Date time.Time

func (j *Date) IsZero() bool {
	return j.Time().IsZero()
}

type DateTime time.Time

func (j *DateTime) IsZero() bool {
	return j.Time().IsZero()
}

type JsonTimestamp time.Time

func (j *JsonTimestamp) IsZero() bool {
	return j.Time().IsZero()
}
