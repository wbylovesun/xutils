package xtime

import "time"

func (w *WeekRange) From() time.Time {
	return w.from
}

func (w *WeekRange) FromAsShortDate() string {
	return ShortDate(w.from)
}

func (w *WeekRange) FromAsTimestamp() int64 {
	return w.from.Unix()
}

func (w *WeekRange) FromAsISOWeek() int {
	return ISOYearWeek(w.from)
}

func (w *WeekRange) FromAsYearMonth() int {
	return YearMonth(w.from)
}

func (w *WeekRange) FromAsQuarter() int {
	return YearQuarter(w.from)
}

func (w *WeekRange) To() time.Time {
	return w.to
}

func (w *WeekRange) ToAsShortDate() string {
	return ShortDate(w.to)
}

func (w *WeekRange) ToAsTimestamp() int64 {
	return w.to.Unix()
}

func (w *WeekRange) ToAsISOWeek() int {
	return ISOYearWeek(w.to)
}

func (w *WeekRange) ToAsYearMonth() int {
	return YearMonth(w.to)
}

func (w *WeekRange) ToAsQuarter() int {
	return YearQuarter(w.to)
}

func ThisWeekRange() WeekRange {
	today := Today()
	ts := StartOfWeek(today)
	te := Today()
	return WeekRange{
		from: ts,
		to:   te,
	}
}

func ThisFullWeekRange() WeekRange {
	today := Today()
	ts := StartOfWeek(today)
	te := EndOfWeek(today)
	return WeekRange{
		from: ts,
		to:   te,
	}
}

func WeekRangeOf(t time.Time) WeekRange {
	ts := StartOfWeek(t)
	te := EndOfWeek(t)
	today := Today()
	if te.After(today) {
		te = today
	}
	return WeekRange{
		from: ts,
		to:   te,
	}
}

func FullWeekRangeOf(t time.Time) WeekRange {
	ts := StartOfWeek(t)
	te := EndOfWeek(t)
	return WeekRange{
		from: ts,
		to:   te,
	}
}

func ISOThisWeekRange() WeekRange {
	today := Today()
	ts := ISOStartOfWeek(today)
	te := Today()
	return WeekRange{
		from: ts,
		to:   te,
	}
}

func ISOThisFullWeekRange() WeekRange {
	today := Today()
	ts := ISOStartOfWeek(today)
	te := ISOEndOfWeek(today)
	return WeekRange{
		from: ts,
		to:   te,
	}
}

func ISOWeekRangeOf(t time.Time) WeekRange {
	ts := ISOStartOfWeek(t)
	te := ISOEndOfWeek(t)
	today := Today()
	if te.After(today) {
		te = today
	}
	return WeekRange{
		from: ts,
		to:   te,
	}
}

func ISOFullWeekRangeOf(t time.Time) WeekRange {
	ts := ISOStartOfWeek(t)
	te := ISOEndOfWeek(t)
	return WeekRange{
		from: ts,
		to:   te,
	}
}
