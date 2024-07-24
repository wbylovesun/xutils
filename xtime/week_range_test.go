package xtime

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestWeekRangeOf(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want WeekRange
	}{
		{
			name: "2024-07-17",
			args: args{
				t: time.Date(2024, 7, 17, 0, 0, 0, 0, time.Local),
			},
			want: WeekRange{
				from: time.Date(2024, 7, 14, 0, 0, 0, 0, time.Local),
				to:   time.Date(2024, 7, 20, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "Today",
			args: args{
				t: Today(),
			},
			want: WeekRange{
				from: StartOfWeek(Today()),
				to:   Today(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeekRangeOf(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeekRangeOf() = %v, want %v", got, tt.want)
				fmt.Println(got.FromAsShortDate(), got.ToAsShortDate(), tt.want.FromAsShortDate(), tt.want.ToAsShortDate())
			}
		})
	}
}

func TestISOWeekRangeOf(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want WeekRange
	}{
		{
			name: "2024-07-17",
			args: args{
				t: time.Date(2024, 7, 17, 0, 0, 0, 0, time.Local),
			},
			want: WeekRange{
				from: time.Date(2024, 7, 15, 0, 0, 0, 0, time.Local),
				to:   time.Date(2024, 7, 21, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "today",
			args: args{
				t: Today(),
			},
			want: WeekRange{
				from: ISOStartOfWeek(time.Now()),
				to:   Today(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ISOWeekRangeOf(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ISOWeekRangeOf() = %v, want %v", got, tt.want)
			}
		})
	}
}
