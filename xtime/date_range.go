package xtime

import (
	"time"
)

func (dr *DateRange) From() time.Time {
	return dr.from
}

func (dr *DateRange) FromAsShortDate() string {
	return ShortDate(dr.from)
}

func (dr *DateRange) FromAsTimestamp() int64 {
	return dr.from.Unix()
}

func (dr *DateRange) FromAsISOWeek() int {
	return ISOYearWeek(dr.from)
}

func (dr *DateRange) FromAsYearMonth() int {
	return YearMonth(dr.from)
}

func (dr *DateRange) FromAsQuarter() int {
	return YearQuarter(dr.from)
}

func (dr *DateRange) To() time.Time {
	return dr.to
}

func (dr *DateRange) ToAsShortDate() string {
	return ShortDate(dr.to)
}

func (dr *DateRange) ToAsTimestamp() int64 {
	return dr.to.Unix()
}

func (dr *DateRange) ToAsISOWeek() int {
	return ISOYearWeek(dr.to)
}

func (dr *DateRange) ToAsYearMonth() int {
	return YearMonth(dr.to)
}

func (dr *DateRange) ToAsQuarter() int {
	return YearQuarter(dr.to)
}

func (dr *DateRange) DayAxis(containsFrom bool, containsTo bool) []string {
	var axis []string
	sd := dr.from
	if !containsFrom {
		sd = dr.from.Add(24 * time.Hour)
	}
	ed := dr.to
	if containsTo {
		ed = dr.to.Add(24 * time.Hour)
	}
	for s := sd; s.Before(ed); s = s.Add(24 * time.Hour) {
		axis = append(axis, ShortDate(s))
	}
	return axis
}
