package xtime

import (
	"reflect"
	"testing"
	"time"
)

func Test_monthGapOf(t *testing.T) {
	type args struct {
		startYear  int
		startMonth int
		endYear    int
		endMonth   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "202201-202206",
			args: args{
				startYear:  2022,
				startMonth: 1,
				endYear:    2022,
				endMonth:   6,
			},
			want: 6,
		},
		{
			name: "202007-202206",
			args: args{
				startYear:  2020,
				startMonth: 7,
				endYear:    2022,
				endMonth:   6,
			},
			want: 24,
		},
		{
			name: "202012-202201",
			args: args{
				startYear:  2020,
				startMonth: 12,
				endYear:    2022,
				endMonth:   1,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monthGapOf(tt.args.startYear, tt.args.startMonth, tt.args.endYear, tt.args.endMonth); got != tt.want {
				t.Errorf("monthGapOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastWeekSameDay(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2024-05-19",
			args: args{
				t: time.Date(2024, 5, 19, 15, 16, 17, 0, time.Local),
			},
			want: time.Date(2024, 5, 12, 15, 16, 17, 0, time.Local),
		},
		{
			name: "2024-03-07",
			args: args{
				t: time.Date(2024, 3, 7, 15, 16, 17, 0, time.Local),
			},
			want: time.Date(2024, 2, 29, 15, 16, 17, 0, time.Local),
		},
		{
			name: "2023-03-07",
			args: args{
				t: time.Date(2023, 3, 7, 15, 16, 17, 0, time.Local),
			},
			want: time.Date(2023, 2, 28, 15, 16, 17, 0, time.Local),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastWeekSameDay(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LastWeekSameDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntDate(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2024-05-19",
			args: args{
				t: time.Date(2024, 5, 19, 15, 16, 17, 0, time.Local),
			},
			want: 20240519,
		},
		{
			name: "2024-01-01",
			args: args{
				t: time.Date(2024, 1, 1, 15, 16, 17, 0, time.Local),
			},
			want: 20240101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DateIntVal(tt.args.t); got != tt.want {
				t.Errorf("DateIntVal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDayOfYear(t *testing.T) {
	type args struct {
		t time.Time
		d int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2024-05-19,20",
			args: args{
				t: time.Date(2024, 5, 19, 15, 16, 17, 0, time.Local),
				d: 20,
			},
			want: time.Date(2024, 1, 20, 15, 16, 17, 0, time.Local),
		},
		{
			name: "2024-05-19,365",
			args: args{
				t: time.Date(2024, 5, 19, 15, 16, 17, 0, time.Local),
				d: 365,
			},
			want: time.Date(2024, 12, 30, 15, 16, 17, 0, time.Local),
		},
		{
			name: "2024-05-19,366",
			args: args{
				t: time.Date(2024, 5, 19, 15, 16, 17, 0, time.Local),
				d: 366,
			},
			want: time.Date(2024, 12, 31, 15, 16, 17, 0, time.Local),
		},
		{
			name: "2024-05-19,367",
			args: args{
				t: time.Date(2024, 5, 19, 15, 16, 17, 0, time.Local),
				d: 367,
			},
			want: time.Date(2025, 1, 1, 15, 16, 17, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeOfDayOfYear(tt.args.t, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDayOfYear1(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2024-05-19",
			args: args{
				t: time.Date(2024, 5, 19, 15, 16, 17, 0, time.Local),
			},
			want: 140,
		},
		{
			name: "2024-12-31",
			args: args{
				t: time.Date(2024, 12, 31, 15, 16, 17, 0, time.Local),
			},
			want: 366,
		},
		{
			name: "2024-01-01",
			args: args{
				t: time.Date(2024, 1, 1, 15, 16, 17, 0, time.Local),
			},
			want: 1,
		},
		{
			name: "2023-12-31",
			args: args{
				t: time.Date(2023, 12, 31, 15, 16, 17, 0, time.Local),
			},
			want: 365,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DayOfYear(tt.args.t); got != tt.want {
				t.Errorf("DayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeeksInYear(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1984-2032",
			args: args{
				start: 1984,
				end:   2032,
			},
			want: []int{52, 52, 52, 52, 53, 52, 52, 52, 52, 52, 53, 52, 52, 52, 52, 52, 53, 52, 52, 52, 52, 53, 52, 52, 52, 52, 52, 53, 52, 52, 52, 52, 53, 52, 52, 52, 52, 52, 53, 52, 52, 52, 52, 52, 53, 52, 52, 52, 52},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			for i := tt.args.start; i <= tt.args.end; i++ {
				got = append(got, WeeksInYear(i))
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeeksInYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstWeekOffset(t *testing.T) {
	type args struct {
		year int
		dow  int
		doy  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2024",
			args: args{
				year: 2024,
				dow:  1,
				doy:  1,
			},
			want: 0,
		},
		{
			name: "2024",
			args: args{
				year: 2024,
				dow:  1,
				doy:  4,
			},
			want: 0,
		},
		{
			name: "2024",
			args: args{
				year: 2024,
				dow:  1,
				doy:  5,
			},
			want: 0,
		},
		{
			name: "2024",
			args: args{
				year: 2024,
				dow:  1,
				doy:  8,
			},
			want: -7,
		},
		{
			name: "2025,1,8",
			args: args{
				year: 2025,
				dow:  1,
				doy:  8,
			},
			want: -2,
		},
		{
			name: "2025,1,1",
			args: args{
				year: 2025,
				dow:  1,
				doy:  1,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstWeekOffset(tt.args.year, tt.args.dow, tt.args.doy); got != tt.want {
				t.Errorf("firstWeekOffset() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		ts string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "202301",
			args: args{
				ts: "202301",
			},
			want:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name: "20231",
			args: args{
				ts: "20231",
			},
			want:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.ts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
