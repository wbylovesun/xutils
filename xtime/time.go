package xtime

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const (
	ShortFormat     = "2006-01-02"
	LongFormat      = "2006-01-02 15:04:05"
	YearMonthFormat = "200601"
	YmdFormat       = "20060102"
	DefaultFormat   = LongFormat
	LongTimeFormat  = "15:04:05"
	ShortTimeFormat = "15:04"
)

var daysBefore = [...]int32{
	0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365,
}

func NewDateRange(start, end time.Time) *DateRange {
	if start.After(end) {
		start, end = end, start
	}
	return &DateRange{
		from: start,
		to:   end,
	}
}

func NewDateRangeByString(start, end, format string) (*DateRange, error) {
	from, err := parseDate(start, format)
	if err != nil {
		return nil, err
	}
	to, err := parseDate(end, format)
	if err != nil {
		return nil, err
	}
	return NewDateRange(from, to), nil
}

func NewComparableDateRange(dr DateRange) *ComparableDateRange {
	return NewComparableDateRangeByTime(dr.from, dr.to)
}

func NewComparableDateRangeByTime(start, end time.Time) *ComparableDateRange {
	return &ComparableDateRange{DateRange{
		from: start,
		to:   end,
	}}
}

func NewComparableDateRangeByString(start, end, format string) (*ComparableDateRange, error) {
	dr, err := NewDateRangeByString(start, end, format)
	if err != nil {
		return nil, err
	}
	return NewComparableDateRange(*dr), nil
}

func NewComparableWeekRange(wr WeekRange) *ComparableWeekRange {
	return &ComparableWeekRange{WeekRange: wr}
}

func NewMonthRange(start, end time.Time) *MonthRange {
	if start.After(end) {
		start, end = end, start
	}
	return &MonthRange{
		from: start,
		to:   end,
	}
}

func NewComparableMonthRangeByTime(start, end time.Time) *ComparableMonthRange {
	return &ComparableMonthRange{MonthRange{
		from: start,
		to:   end,
	}}
}

func YearMonth(t time.Time) int {
	toMonth, _ := strconv.Atoi(t.Format(YearMonthFormat))
	return toMonth
}

func YearQuarter(t time.Time) int {
	return t.Year()*100 + int(math.Ceil(float64(t.Month())/3))
}

func ShortDate(t time.Time) string {
	return t.Format(ShortFormat)
}

func LongDate(t time.Time) string {
	return t.Format(LongFormat)
}

func YmdDate(t time.Time) string {
	return t.Format(YmdFormat)
}

func LongTime(t time.Time) string {
	return t.Format(LongTimeFormat)
}

func ShortTime(t time.Time) string {
	return t.Format(ShortTimeFormat)
}

func WithTime(t time.Time, hour, minute, second int) time.Time {
	return time.Date(
		t.Year(), t.Month(), t.Day(),
		hour, minute, second, 0, time.Local,
	)
}

func FirstDay(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		1,
		0, 0, 0, 0, time.Local,
	)
}

func LastDay(t time.Time) time.Time {
	daysIn := DaysIn(t)
	return time.Date(
		t.Year(),
		t.Month(),
		daysIn,
		0, 0, 0, 0, time.Local,
	)
}

func AddDays(t time.Time, days int) time.Time {
	duration := time.Duration(days) * 24 * time.Hour
	var t2 time.Time
	if t.IsZero() {
		t2 = time.Now()
	} else {
		t2 = t
	}
	return t2.Add(duration)
}

func SubtractDays(t time.Time, days int) time.Time {
	return AddDays(t, -days)
}

func IsLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

func DaysIn(t time.Time) int {
	return DaysInYearMonth(t.Year(), int(t.Month()))
}

func DaysInYearMonth(y, m int) int {
	if IsLeap(y) && m == int(time.February) {
		return 29
	}
	return int(daysBefore[m] - daysBefore[m-1])
}

func LastSecondOf(t time.Time) time.Time {
	return WithTime(t, 23, 59, 59)
}

func DateIntVal(t time.Time) int {
	return t.Year()*10000 + int(t.Month())*100 + t.Day()
}

func ParseDateIntVal(d int) (time.Time, error) {
	ds := strconv.Itoa(d)
	return Parse(ds)
}

// DayOfYear 年积日，获取一年中的第几天，从1开始
func DayOfYear(t time.Time) int {
	nt := time.Date(t.Year(), 1, 0, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	gap := int(t.Sub(nt).Hours())
	gap /= 24
	return gap
}

// TimeOfDayOfYear 设置一年中的第几天，从1开始
// 接受从 1 到 366 的数字。如果超出范围，它将冒泡到年份。
func TimeOfDayOfYear(t time.Time, d int) time.Time {
	nt := time.Date(t.Year(), time.Month(1), 0, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	nt = nt.AddDate(0, 0, d)
	return nt
}

// WeeksInYear 获取一年中的周数
// 参考：moment.js
func WeeksInYear(year int) int {
	dayOfWeek := 0
	dayOfYear := 6
	daysInYear := 365
	if IsLeap(year) {
		daysInYear = 366
	}
	weekOffset := firstWeekOffset(year, int(dayOfWeek), dayOfYear)
	weekOffsetNext := firstWeekOffset(year+1, int(dayOfWeek), dayOfYear)
	return (daysInYear - weekOffset + weekOffsetNext) / 7
}

func firstWeekOffset(year, dow, doy int) int {
	fwd := 7 + dow - doy
	fwdlw := (7 + int(time.Date(year, 1, fwd, 0, 0, 0, 0, time.UTC).Weekday()) - dow) % 7
	return -fwdlw + fwd - 1
}

func parseDate(date, format string) (time.Time, error) {
	if format == "" {
		format = ShortFormat
	}
	return time.ParseInLocation(format, date, time.Local)
}

var formatMap = map[byte][]byte{
	// Year
	'Y': []byte("2006"), // long year. eg: 1999, 2003
	'y': []byte("06"),   // short year. eg: 99 or 03
	// Month
	'm': {'0', '1'},        // 01 through 12
	'n': {'1'},             // 1 through 12
	'M': []byte("Jan"),     // Jan through Dec
	'F': []byte("January"), // January through December(full month, php)
	// Day
	'j': {'2'},            // day of the month, 1 to 31
	'd': []byte("02"),     // day of the month, 01 to 31
	'D': []byte("Mon"),    // weekday. Mon through Sun(php)
	'w': []byte("Mon"),    // weekday. Mon through Sun
	'W': []byte("Monday"), // long weekday. Sunday through Saturday
	'l': []byte("Monday"), // long weekday. Sunday through Saturday(php)
	'z': []byte("002"),    // day of the year, 0 through 365
	// Hour
	'H': []byte("15"), // 00 through 23
	'h': []byte("03"), // 01 through 12
	'g': {'3'},        // 1 through 12
	'G': []byte("15"), // go not support 0-23, use 00-23 instead
	// Minutes - 'i' is second char of 'minutes'
	'I': []byte("04"), // 00 to 59
	'i': []byte("4"),  // 0 to 59
	// Seconds
	'S': []byte("05"), // 00 to 59
	's': []byte("5"),  // 0 to 59
	// Time
	'a': []byte("pm"),      // am or pm
	'A': []byte("PM"),      // AM or PM
	'v': []byte(".000"),    // Milliseconds eg: 654
	'u': []byte(".000000"), // Microseconds eg: 654321
	// Timezone
	'e': []byte("MST"),    // Timezone identifier. eg: UTC, GMT, Atlantic/Azores
	'Z': []byte("Z07"),    // Timezone abbreviation, if known; otherwise the GMT offset. Examples: EST, MDT, +05
	'O': []byte("Z0700"),  // Difference to Greenwich time (GMT) without colon between hours and minutes. Example: +0200
	'P': []byte("Z07:00"), // Difference to Greenwich time (GMT) with colon between hours and minutes. Example: +02:00
	// Full Date/Time
	'c': []byte(time.RFC3339),  // ISO 8601 date. eg: 2004-02-12T15:19:21+00:00
	'r': []byte(time.RFC1123Z), // » RFC 2822/» RFC 5322 formatted date. eg: Thu, 21 Dec 2000 16:01:07 +0200
}

// BuildFormat convert chars date template to Go date layout.
func BuildFormat(template string) string {
	if template == "" {
		return DefaultFormat
	}

	bytes := []byte(template)
	bs := make([]byte, 0, 24)
	for _, c := range bytes {
		if b, ok := formatMap[c]; ok {
			bs = append(bs, b...)
		} else {
			bs = append(bs, c)
		}
	}

	return string(bs)
}

func MustOf(ts string) time.Time {
	t, err := Parse(ts)
	if err != nil {
		return time.Time{}
	}
	return t
}

// MustOfDateIntVal 将Ymd格式的日期数值转换为time.Time
func MustOfDateIntVal(ts int) time.Time {
	return MustOf(strconv.Itoa(ts))
}

func Parse(ts string) (time.Time, error) {
	var layouts = []string{
		"2006-01-02",
		"2006/01/02",
		"2006.01.02",
		"20060102",
		"2006/01/02 15:04:05",
		"2006-01-02 15:04:05",
		"2006.01.02 15:04:05",
		"20060102150405",
		"2006/01/02 15:04",
		"2006-01-02 15:04",
		"2006.01.02 15:04",
		"200601021504",
		"2006/01/02 15",
		"2006-01-02 15",
		"2006.01.02 15",
		"2006010215",
		"2006/01",
		"2006-01",
		"20060.01",
		"200601",
		"2006/1",
		"2006-1",
		"2006.1",
		"20061",
		"2006",
		"01/02",
		"01-02",
		"01.02",
		"0102",
		"15:04:05",
		"15:04",
		"1504",
		time.RFC3339,
		time.RFC822,
		time.RFC822Z,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
	}
	for _, layout := range layouts {
		t, err := time.ParseInLocation(layout, ts, time.Local)
		if nil == err && !t.IsZero() {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("cannot parse value %s", ts)
}

func monthGap(start, end time.Time) int {
	return monthGapOf(
		start.Year(),
		int(start.Month()),
		end.Year(),
		int(end.Month()),
	)
}

func monthGapOf(startYear, startMonth, endYear, endMonth int) int {
	gap := 0
	for i := startYear; i <= endYear; i++ {
		for j := 1; j <= 12; j++ {
			if i == startYear && j < startMonth {
				continue
			}
			if i == endYear && j > endMonth {
				break
			}
			gap++
		}
	}
	return gap
}
