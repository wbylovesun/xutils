package main

import (
	"fmt"
	"github.com/wbylovesun/xutils/xcalc"
	"github.com/wbylovesun/xutils/xtime"
)

type item struct {
	Count  int
	Amount int
	Ticket int
}

func main() {
	sequence()
}

func sequence() {
	ts, _ := xtime.NewTimeSpan(5)
	segs := ts.RangeSegs("09:00", "16:00")
	fmt.Println(segs)
	m := map[string]int{
		"09:05": 20,
		"09:30": 28,
		"09:35": 33,
		"09:40": 34,
		"09:45": 44,
		"10:00": 69,
		"14:45": 80,
		"15:30": 88,
	}
	fmt.Println(xcalc.SequenceValuePad(segs, m, func(last int, k int) int {
		return 0
	}))
	fmt.Println(xcalc.SequenceValuePad(segs, m, func(last int, k int) int {
		if k == 0 {
			return 0
		}
		return last
	}))
	m2 := map[string]*item{
		"09:05": {
			Count:  1,
			Amount: 500,
			Ticket: 2,
		},
		"09:40": {
			Count:  5,
			Amount: 2500,
			Ticket: 5,
		},
		"10:50": {
			Count:  10,
			Amount: 5000,
			Ticket: 8,
		},
	}
	//m3 := xcalc.SequenceValuePad(segs, m2, )
	//for _, k := range segs {
	//	v := m3[k]
	//	fmt.Println(k, v.Count, v.Amount, v.Ticket)
	//}
	m3 := xcalc.IncrementalSequenceValueDiff(segs, m2, func(current, last *item) *item {
		return &item{
			Count:  current.Count - last.Count,
			Amount: current.Amount - last.Amount,
			Ticket: current.Ticket - last.Ticket,
		}
	}, nil)
	fmt.Println(m3)

	m4 := xcalc.IncrementalSequenceValueDiff(segs, m2, func(current, last *item) *item {
		return &item{
			Count:  current.Count - last.Count,
			Amount: current.Amount - last.Amount,
			Ticket: current.Ticket - last.Ticket,
		}
	}, func(v *item, k int) *item {
		if k == 0 {
			return &item{
				Count:  0,
				Amount: 0,
				Ticket: 0,
			}
		}
		return &item{
			Count:  v.Count,
			Amount: v.Amount,
			Ticket: v.Ticket,
		}
	})
	for _, k := range segs {
		v := m4[k]
		fmt.Println(k, v.Count, v.Amount, v.Ticket)
	}
}
