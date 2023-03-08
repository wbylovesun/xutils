package xtime

import (
	"fmt"
	"github.com/wbylovesun/xutils"
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

// PointOfTimeSlice 指的是按指定时长切片后形成的时间点，如09:00, 12:15, 13:20
type PointOfTimeSlice string

func (ts *timeSpan) init() {
	for i := 0; i <= 1440; i += ts.span {
		ts.timeSpans = append(ts.timeSpans, i)
		ts.timeSegs = append(ts.timeSegs, ts.toHourMinute(i))
	}
}

func (ts *timeSpan) toHourMinute(t int) string {
	hour := t / 60
	minute := t % 60
	return fmt.Sprintf("%02d:%02d", hour, minute)
}

// Index 从0开始的seg下标,i值为hour*60+minute
func (ts *timeSpan) Index(i int) int {
	return i / ts.span
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

// Of
//
// Deprecated: 请使用OfNextSeg替代
func (ts *timeSpan) Of(i int) string {
	x := ts.Index(i) + 1
	return ts.timeSegs[x]
}

func (ts *timeSpan) OfNextSeg(i int) string {
	x := ts.Index(i) + 1
	return ts.timeSegs[x]
}

// OfTime 整切片值的时间会被归到下一个时间切片
//
// 如5分钟切片，09:00会归为09:05，09:10会归为09:15
//
// Deprecated: 请使用AlignToNextSeg替代
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

// AlignToNextSeg OfTime的别名，用于向后对齐，如5分钟的切片间隔，16:26对齐到16:30
func (ts *timeSpan) AlignToNextSeg(i string) (string, error) {
	if !ts.isValidHourMinute(i) {
		return "", fmt.Errorf("invalid time format: HH:MM")
	}
	parts := strings.Split(i, ":")
	hour, _ := strconv.Atoi(parts[0])
	minute, _ := strconv.Atoi(parts[1])
	n := hour*60 + minute
	return ts.OfNextSeg(n), nil
}

// AlignToPrevSeg 用于向前对齐，如5分钟的切片间隔，16:26对齐到16:25
func (ts *timeSpan) AlignToPrevSeg(i string) (string, error) {
	if !ts.isValidHourMinute(i) {
		return "", fmt.Errorf("invalid time format: HH:MM")
	}
	parts := strings.Split(i, ":")
	hour, _ := strconv.Atoi(parts[0])
	minute, _ := strconv.Atoi(parts[1])
	n := hour*60 + minute
	return ts.OfPrevSeg(n), nil
}

func (ts *timeSpan) TimeOfNextAlignSeg(i time.Time) (time.Time, error) {
	seg, err := ts.AlignToNextSeg(i.Format("15:04"))
	if err != nil {
		return time.Time{}, err
	}
	t, err := Parse(fmt.Sprintf("%s %s", ShortDate(i), seg))
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func (ts *timeSpan) TimeOfPrevAlignSeg(i time.Time) (time.Time, error) {
	seg, err := ts.AlignToPrevSeg(i.Format("15:04"))
	if err != nil {
		return time.Time{}, err
	}
	t, err := Parse(fmt.Sprintf("%s %s", ShortDate(i), seg))
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func (ts *timeSpan) OfPrevSeg(i int) string {
	x := ts.Index(i)
	return ts.timeSegs[x]
}

func (ts *timeSpan) Segs() []string {
	return ts.timeSegs
}

func (ts *timeSpan) TodaySegs() []string {
	segs := ts.timeSegs
	t, _ := ts.AlignToNextSeg(time.Now().Format("15:04"))
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

func (ts *timeSpan) Fill(m map[PointOfTimeSlice]int, pots ...PointOfTimeSlice) map[PointOfTimeSlice]int {
	deepCopy := xutils.Copy(m)
	mCopy := deepCopy.(map[PointOfTimeSlice]int)
	var till int
	if len(pots) == 0 || !ts.isValidHourMinute(string(pots[0])) {
		till, _ = ts.timeValue(time.Now().Format("15:04"))
	} else {
		till, _ = ts.timeValue(string(pots[0]))
	}
	tillHour := till / 60
	tillMinute := till % 60
	lastVal := 0
	for i := 0; i <= 1440; i += ts.span {
		hour := i / 60
		minute := i % 60
		if hour > tillHour {
			break
		}
		if hour == tillHour && minute > tillMinute {
			break
		}
		sliceTime := fmt.Sprintf("%02d:%02d", hour, minute)
		if count, ok := mCopy[PointOfTimeSlice(sliceTime)]; ok {
			lastVal = count
		} else {
			mCopy[PointOfTimeSlice(sliceTime)] = lastVal
		}
	}
	return mCopy
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
	spans := []int{5, 10, 15, 20, 30, 60, 90, 120, 180, 240, 360, 480, 720}
	if !xslice.Contains(spans, span) {
		return nil, fmt.Errorf("invalid span specified")
	}
	ts := new(timeSpan)
	ts.span = span
	ts.init()
	return ts, nil
}
