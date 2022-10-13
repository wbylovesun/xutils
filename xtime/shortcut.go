package xtime

import "time"

func ThisYearRange(includeToday ...bool) *DateRange {
	include := false
	if len(includeToday) > 0 {
		include = includeToday[0]
	}
	now := time.Now()
	from := time.Date(now.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
	to := WithTime(now.Add(-24*time.Hour), 0, 0, 0)
	if include {
		to = WithTime(now, 0, 0, 0)
	}
	return NewDateRange(
		from,
		to,
	)
}

func LastYearRange() *DateRange {
	now := time.Now()
	from := time.Date(now.Year()-1, time.January, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(from.Year(), time.December, 31, 0, 0, 0, 0, time.Local)
	return NewDateRange(
		from,
		to,
	)
}

func Today() time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func TodayLongDate() string {
	return LongDate(time.Now())
}

func TodayShortDate() string {
	return ShortDate(time.Now())
}

func TodayYmdDate() string {
	return YmdDate(time.Now())
}

func Yesterday() time.Time {
	return Today().AddDate(0, 0, -1)
}

func YesterdayLongDate() string {
	return LongDate(time.Now().AddDate(0, 0, -1))
}

func YesterdayShortDate() string {
	return ShortDate(time.Now().AddDate(0, 0, -1))
}

func YesterdayYmdDate() string {
	return YmdDate(time.Now().AddDate(0, 0, -1))
}

func Tomorrow() time.Time {
	return Today().AddDate(0, 0, 1)
}

func TomorrowLongDate() string {
	return LongDate(time.Now().AddDate(0, 0, 1))
}

func TomorrowShortDate() string {
	return ShortDate(time.Now().AddDate(0, 0, 1))
}

func TomorrowYmdDate() string {
	return YmdDate(time.Now().AddDate(0, 0, 1))
}

func ISOLastWeek() int {
	return ISOYearWeek(time.Now().AddDate(0, 0, -7))
}

func FirstDayOfThisMonth() time.Time {
	now := time.Now()
	return time.Date(
		now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local,
	)
}

func LastDayOfThisMonth() time.Time {
	now := time.Now()
	daysIn := DaysIn(now)
	return time.Date(
		now.Year(), now.Month(), daysIn, 0, 0, 0, 0, time.Local,
	)
}

// Latest30Days 近30天，包含今天
func Latest30Days() *DateRange {
	return LatestNDays(30)
}

//Passed30Days 过去30天，不包含今天
func Passed30Days() *DateRange {
	return PassedNDays(30)
}

// LatestNDays 最近N天，包含今天
func LatestNDays(n int) *DateRange {
	today := Today()
	from := today.AddDate(0, 0, -1*n+1)
	return NewDateRange(
		from,
		today,
	)
}

// PassedNDays 过去N天，不包含今天
func PassedNDays(n int) *DateRange {
	yesterday := Yesterday()
	from := yesterday.AddDate(0, 0, -1*n+1)
	return NewDateRange(from, yesterday)
}

// Lastest1Year 最近1年，包含今天
func Lastest1Year() *DateRange {
	today := Today()
	from := today.AddDate(-1, 0, 0)
	return NewDateRange(from, today)
}

// Passed1Year 过去1年，不包含今天
func Passed1Year() *DateRange {
	yesterday := Yesterday()
	from := yesterday.AddDate(-1, 0, 0)
	return NewDateRange(from, yesterday)
}

// Latest12Months 最近12个月，包含当月
func Latest12Months() *MonthRange {
	return LatestNMonths(12)
}

// Passed12Months 过去12个月，不包含当月
func Passed12Months() *MonthRange {
	return PassedNMonths(12)
}

func LatestNMonths(n int) *MonthRange {
	today := Today()
	to := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.Local)
	from := to.AddDate(0, -1*n+1, 0)
	return &MonthRange{
		from: from,
		to:   to,
	}
}

func PassedNMonths(n int) *MonthRange {
	today := Today()
	to := time.Date(today.Year(), today.Month()-1, 1, 0, 0, 0, 0, time.Local)
	from := to.AddDate(0, -1*n+1, 0)
	return &MonthRange{
		from: from,
		to:   to,
	}
}
