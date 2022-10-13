package xtime

// MonthOnMonth 环比
func (cmr *ComparableMonthRange) MonthOnMonth() *ComparableMonthRange {
	gap := monthGap(cmr.from, cmr.to)
	momTo := cmr.from.AddDate(0, -1, 0)
	momFrom := momTo.AddDate(0, -1*gap+1, 0)
	return &ComparableMonthRange{
		MonthRange{
			from: momFrom,
			to:   momTo,
		},
	}
}

//YearOnYear 年同比
func (cmr ComparableMonthRange) YearOnYear() *ComparableMonthRange {
	yoyFrom := cmr.from.AddDate(-1, 0, 0)
	yoyTo := cmr.to.AddDate(-1, 0, 0)
	return &ComparableMonthRange{
		MonthRange{
			from: yoyFrom,
			to:   yoyTo,
		},
	}
}
