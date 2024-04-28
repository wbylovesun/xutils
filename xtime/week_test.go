package xtime

import (
	"reflect"
	"testing"
	"time"
)

func TestStartOfWeek(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2024-04-24",
			args: args{
				t: time.Date(2024, 4, 24, 15, 30, 0, 0, time.Local),
			},
			want: time.Date(2024, 4, 21, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2024-04-27",
			args: args{
				t: time.Date(2024, 4, 27, 15, 30, 0, 0, time.Local),
			},
			want: time.Date(2024, 4, 21, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2024-04-21",
			args: args{
				t: time.Date(2024, 4, 21, 15, 30, 0, 0, time.Local),
			},
			want: time.Date(2024, 4, 21, 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartOfWeek(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StartOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndOfWeek(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2024-04-24",
			args: args{
				t: time.Date(2024, 4, 24, 15, 30, 0, 0, time.Local),
			},
			want: time.Date(2024, 4, 27, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2024-04-27",
			args: args{
				t: time.Date(2024, 4, 27, 15, 30, 0, 0, time.Local),
			},
			want: time.Date(2024, 4, 27, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2024-04-21",
			args: args{
				t: time.Date(2024, 4, 21, 15, 30, 0, 0, time.Local),
			},
			want: time.Date(2024, 4, 27, 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndOfWeek(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EndOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeekDateRange(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want DateRange
	}{
		{
			name: "2024-04-24",
			args: args{
				t: time.Date(2024, 4, 24, 15, 30, 0, 0, time.Local),
			},
			want: DateRange{
				from: time.Date(2024, 4, 21, 0, 0, 0, 0, time.Local),
				to:   time.Date(2024, 4, 27, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "2024-04-21",
			args: args{
				t: time.Date(2024, 4, 21, 15, 30, 0, 0, time.Local),
			},
			want: DateRange{
				from: time.Date(2024, 4, 21, 0, 0, 0, 0, time.Local),
				to:   time.Date(2024, 4, 27, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "2024-04-27",
			args: args{
				t: time.Date(2024, 4, 27, 15, 30, 0, 0, time.Local),
			},
			want: DateRange{
				from: time.Date(2024, 4, 21, 0, 0, 0, 0, time.Local),
				to:   time.Date(2024, 4, 27, 0, 0, 0, 0, time.Local),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WeekDateRange(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeekDateRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
