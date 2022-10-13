package xtime

import (
	"strconv"
	"time"
)

func (mr *MonthRange) From() time.Time {
	return mr.from
}

func (mr *MonthRange) FromYearMonth() int {
	return mr.from.Year()*100 + int(mr.from.Month())
}

func (mr *MonthRange) FromYearMonthAsString() string {
	return strconv.Itoa(mr.FromYearMonth())
}

func (mr *MonthRange) FirstDayOfFrom() time.Time {
	return FirstDay(mr.from)
}

func (mr *MonthRange) LastDayOfFrom() time.Time {
	return LastDay(mr.from)
}

func (mr *MonthRange) To() time.Time {
	return mr.to
}

func (mr *MonthRange) ToYearMonth() int {
	return mr.to.Year()*100 + int(mr.to.Month())
}

func (mr *MonthRange) ToYearMonthAsString() string {
	return strconv.Itoa(mr.ToYearMonth())
}

func (mr *MonthRange) FirstDayOfTo() time.Time {
	return FirstDay(mr.to)
}

func (mr *MonthRange) LastDayOfTo() time.Time {
	return LastDay(mr.to)
}

func (mr *MonthRange) Slice() []int {
	var slices []int
	fym := mr.FromYearMonth()
	tym := mr.ToYearMonth()
	for i := fym; i <= tym; i++ {
		y := i / 100
		m := i % 100
		if m > 12 {
			y = y + 1
			m = 1
			i = y*100 + m
		}
		slices = append(slices, y*100+m)
	}
	return slices
}
