package xtime

import (
	"fmt"
	"github.com/wbylovesun/xutils/xslice"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type timeSpan struct {
	span      int
	timeSpans []int
	timeSegs  []string
}

func (ts *timeSpan) init() {
	segs := 60 / ts.span
	l := 24 * segs * ts.span
	for i := 0; i <= l; i += ts.span {
		ts.timeSpans = append(ts.timeSpans, i)
		ts.timeSegs = append(ts.timeSegs, ts.toHourMinute(i))
	}
}

func (ts *timeSpan) toHourMinute(t int) string {
	hour := t / 60
	minute := t % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

func (ts *timeSpan) Index(i int) int {
	x := i / ts.span
	return x + 1
}

func (ts *timeSpan) IndexByTime(i string) (int, error) {
	if !ts.isValidHourMinute(i) {
		return 0, fmt.Errorf("invalid time format: HH:MM")
	}
	parts := strings.Split(i, ":")
	hour, _ := strconv.Atoi(parts[0])
	minute, _ := strconv.Atoi(parts[1])
	n := hour*60 + minute
	return ts.Index(n), nil
}

func (ts *timeSpan) Of(i int) string {
	x := ts.Index(i)
	return ts.timeSegs[x]
}

// OfTime 整切片值的时间会被归到下一个时间切片
//
// 如5分钟切片，09:00会归为09:05，09:10会归为09:15
func (ts *timeSpan) OfTime(i string) (string, error) {
	if !ts.isValidHourMinute(i) {
		return "", fmt.Errorf("invalid time format: HH:MM")
	}
	parts := strings.Split(i, ":")
	hour, _ := strconv.Atoi(parts[0])
	minute, _ := strconv.Atoi(parts[1])
	n := hour*60 + minute
	return ts.Of(n), nil
}

func (ts *timeSpan) Segs() []string {
	return ts.timeSegs
}

func (ts *timeSpan) TodaySegs() []string {
	segs := ts.timeSegs
	t, _ := ts.OfTime(time.Now().Format("15:04"))
	var todaySegs []string
	for _, seg := range segs {
		if seg > t {
			break
		}
		todaySegs = append(todaySegs, seg)
	}
	return todaySegs
}

// RangeSegs 以给定的from, to（包含）生成时间切片范围
//
// 如from=09:00, to=16:00，以30分钟为span，生成[09:00, 09:30, ..., 15:30, 16:00]
func (ts *timeSpan) RangeSegs(from, to string) []string {
	segs := ts.timeSegs
	ftv, _ := ts.timeValue(from)
	ttv, _ := ts.timeValue(to)
	ft := ts.Of(ftv - 1)
	tt := ts.Of(ttv - 1)
	var rangeSegs []string
	for _, seg := range segs {
		if seg < ft {
			continue
		}
		if seg > tt {
			break
		}
		rangeSegs = append(rangeSegs, seg)
	}
	return rangeSegs
}

func (ts *timeSpan) isValidHourMinute(i string) bool {
	ptn, err := regexp.Compile(`^(\d|[01]\d|2[0-4]):(\d|[0-5]\d)$`)
	if err != nil {
		return false
	}
	if ptn.MatchString(i) {
		return true
	}
	return false
}

func (ts *timeSpan) timeValue(i string) (int, error) {
	if !ts.isValidHourMinute(i) {
		return 0, fmt.Errorf("invalid time format: HH:MM")
	}
	parts := strings.Split(i, ":")
	hour, _ := strconv.Atoi(parts[0])
	minute, _ := strconv.Atoi(parts[1])
	n := hour*60 + minute
	return n, nil
}

func NewTimeSpan(span int) (*timeSpan, error) {
	spans := []int{5, 10, 15, 20, 30, 60}
	if !xslice.Contains(spans, span) {
		return nil, fmt.Errorf("invalid span specified")
	}
	ts := new(timeSpan)
	ts.span = span
	ts.init()
	return ts, nil
}
