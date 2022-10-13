package xtime

import "time"

// PeriodOnPeriod 非整月的周期环比
func (cdr *ComparableDateRange) PeriodOnPeriod() *ComparableDateRange {
	duration := cdr.to.Sub(cdr.from)
	pe := cdr.from.Add(-24 * time.Hour)
	ps := pe.Add(-1 * duration)
	return &ComparableDateRange{
		DateRange: DateRange{
			from: ps,
			to:   pe,
		},
	}
}

// MonthOnMonth 环比
func (cdr *ComparableDateRange) MonthOnMonth() *ComparableDateRange {
	momFrom := cdr.from.AddDate(0, -1, 0)
	momTo := cdr.to.AddDate(0, -1, 0)
	return &ComparableDateRange{
		DateRange: DateRange{
			from: momFrom,
			to:   momTo,
		},
	}
}

//YearOnYear 年同比
func (cdr ComparableDateRange) YearOnYear() *ComparableDateRange {
	yoyFrom := cdr.from.AddDate(-1, 0, 0)
	yoyTo := cdr.to.AddDate(-1, 0, 0)
	return &ComparableDateRange{
		DateRange: DateRange{
			from: yoyFrom,
			to:   yoyTo,
		},
	}
}
