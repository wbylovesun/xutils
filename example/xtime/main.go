package main

import (
	"fmt"
	"github.com/wbylovesun/xutils/xtime"
)

func main() {
	m := map[xtime.PointOfTimeSlice]int{
		"02:00": 10,
		"06:00": 30,
		"10:00": 60,
		"16:00": 100,
	}
	ts, _ := xtime.NewTimeSpan(120)
	m2 := ts.Fill(m, "16:26")
	fmt.Println("till 16:26", m2)

	m3 := ts.Fill(m, "22:00")
	fmt.Println("till 22:00", m3)

	m4 := ts.Fill(m, "23:20")
	fmt.Println("till 23:20", m4)

	m5 := ts.Fill(m)
	fmt.Println("till Now:", m5)

	ts2, _ := xtime.NewTimeSpan(5)
	m6 := ts2.Fill(m, "16:26")
	fmt.Println("m6:", m6)

	fmt.Println(ts2.AlignToPrevSeg("16:26"))
	idx, _ := ts2.IndexByTime("00:06")
	idx2, _ := ts2.IndexByTime("16:26")
	fmt.Println(ts2.Index(0), ts2.Index(5), ts2.Index(6), ts2.Index(16*60+26), idx, idx2)

	t, _ := xtime.Parse("2023-03-08 16:26")
	fmt.Println(ts2.TimeOfNextAlignSeg(t))
	fmt.Println(ts2.TimeOfPrevAlignSeg(t))

	w := xtime.ISOThisWeekRange()
	fmt.Println(w.FromAsShortDate(), w.ToAsShortDate())
	c := xtime.NewComparableWeekRange(w)
	nw := c.LastWeekRange()
	fmt.Println("nw=", nw.FromAsShortDate(), nw.ToAsShortDate())
	nw2 := c.SameDurationLastWeekRange()
	fmt.Println("nw2=", nw2.FromAsShortDate(), nw2.ToAsShortDate())
	nw3 := c.NWeeksAgoWeekRange(2)
	fmt.Println("nw3=", nw3.FromAsShortDate(), nw3.ToAsShortDate())
	nws1 := c.LastNWeeksWeekRange(2)
	for _, nw4 := range nws1 {
		fmt.Println("nw4=", nw4.FromAsShortDate(), nw4.ToAsShortDate())
	}
	nws2 := c.SameDurationLastNWeeksWeekRange(2)
	for _, nw5 := range nws2 {
		fmt.Println("nw5=", nw5.FromAsShortDate(), nw5.ToAsShortDate())
	}
}
