package xtime

import (
	"reflect"
	"testing"
	"time"
)

func TestISOWeekdayTime(t *testing.T) {
	type args struct {
		t       time.Time
		weekday int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2024-04-24,7",
			args: args{
				t:       time.Date(2024, 4, 24, 15, 30, 0, 0, time.Local),
				weekday: 7,
			},
			want: time.Date(2024, 4, 28, 15, 30, 0, 0, time.Local),
		},
		{
			name: "2024-04-24,4",
			args: args{
				t:       time.Date(2024, 4, 24, 15, 30, 0, 0, time.Local),
				weekday: 4,
			},
			want: time.Date(2024, 4, 25, 15, 30, 0, 0, time.Local),
		},
		{
			name: "2024-04-24,3",
			args: args{
				t:       time.Date(2024, 4, 24, 15, 30, 0, 0, time.Local),
				weekday: 3,
			},
			want: time.Date(2024, 4, 24, 15, 30, 0, 0, time.Local),
		},
		{
			name: "2024-04-24,2",
			args: args{
				t:       time.Date(2024, 4, 24, 15, 30, 0, 0, time.Local),
				weekday: 2,
			},
			want: time.Date(2024, 4, 23, 15, 30, 0, 0, time.Local),
		},
		{
			name: "2024-04-24,1",
			args: args{
				t:       time.Date(2024, 4, 24, 15, 30, 0, 0, time.Local),
				weekday: 1,
			},
			want: time.Date(2024, 4, 22, 15, 30, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeOfISOWeekday(tt.args.t, tt.args.weekday); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeOfISOWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISOWeekday(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2024-04-22",
			args: args{
				t: time.Date(2024, 4, 22, 15, 30, 0, 0, time.Local),
			},
			want: 1,
		},
		{
			name: "2024-04-28",
			args: args{
				t: time.Date(2024, 4, 28, 15, 30, 0, 0, time.Local),
			},
			want: 7,
		},
		{
			name: "2023-12-28",
			args: args{
				t: time.Date(2023, 12, 28, 15, 30, 0, 0, time.Local),
			},
			want: 4,
		},
		{
			name: "2022-12-31",
			args: args{
				t: time.Date(2022, 12, 31, 15, 30, 0, 0, time.Local),
			},
			want: 6,
		},
		{
			name: "2023-01-01",
			args: args{
				t: time.Date(2023, 1, 1, 15, 30, 0, 0, time.Local),
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ISOWeekday(tt.args.t); got != tt.want {
				t.Errorf("ISOWeekday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISOWeeksInYear(t *testing.T) {
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
			want: []int{52, 52, 52, 53, 52, 52, 52, 52, 53, 52, 52, 52, 52, 52, 53, 52, 52, 52, 52, 52, 53, 52, 52, 52, 52, 53, 52, 52, 52, 52, 52, 53, 52, 52, 52, 52, 53, 52, 52, 52, 52, 52, 53, 52, 52, 52, 52, 52, 53},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []int
			for year := tt.args.start; year <= tt.args.end; year++ {
				got = append(got, ISOWeeksInYear(year))
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ISOWeeksInYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISOWeekDateRange(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want DateRange
	}{
		{
			name: "2024-04-27",
			args: args{
				t: time.Date(2024, 4, 27, 15, 16, 17, 0, time.Local),
			},
			want: DateRange{
				from: time.Date(2024, 4, 22, 0, 0, 0, 0, time.Local),
				to:   time.Date(2024, 4, 28, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "2024-04-28",
			args: args{
				t: time.Date(2024, 4, 28, 15, 16, 17, 0, time.Local),
			},
			want: DateRange{
				from: time.Date(2024, 4, 22, 0, 0, 0, 0, time.Local),
				to:   time.Date(2024, 4, 28, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "2024-04-22",
			args: args{
				t: time.Date(2024, 4, 22, 15, 16, 17, 0, time.Local),
			},
			want: DateRange{
				from: time.Date(2024, 4, 22, 0, 0, 0, 0, time.Local),
				to:   time.Date(2024, 4, 28, 0, 0, 0, 0, time.Local),
			},
		},
		{
			name: "2024-04-30",
			args: args{
				t: time.Date(2024, 4, 30, 15, 16, 17, 0, time.Local),
			},
			want: DateRange{
				from: time.Date(2024, 4, 29, 0, 0, 0, 0, time.Local),
				to:   time.Date(2024, 5, 5, 0, 0, 0, 0, time.Local),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ISOWeekDateRange(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WeekDateRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISOStartOfWeek(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2024-04-27",
			args: args{
				t: time.Date(2024, 4, 27, 15, 16, 17, 0, time.Local),
			},
			want: time.Date(2024, 4, 22, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2024-04-28",
			args: args{
				t: time.Date(2024, 4, 28, 15, 16, 17, 0, time.Local),
			},
			want: time.Date(2024, 4, 22, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2024-04-22",
			args: args{
				t: time.Date(2024, 4, 22, 15, 16, 17, 0, time.Local),
			},
			want: time.Date(2024, 4, 22, 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ISOStartOfWeek(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ISOStartOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISOEndOfWeek(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2024-04-27",
			args: args{
				t: time.Date(2024, 4, 27, 15, 16, 17, 0, time.Local),
			},
			want: time.Date(2024, 4, 28, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2024-04-28",
			args: args{
				t: time.Date(2024, 4, 28, 15, 16, 17, 0, time.Local),
			},
			want: time.Date(2024, 4, 28, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2024-04-22",
			args: args{
				t: time.Date(2024, 4, 22, 15, 16, 17, 0, time.Local),
			},
			want: time.Date(2024, 4, 28, 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ISOEndOfWeek(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ISOEndOfWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISOInfo(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "2024-01-30",
			args: args{
				t: time.Date(2024, 1, 30, 15, 16, 17, 0, time.Local),
			},
			want: "2024-W05-2",
		},
		{
			name: "2024-04-28",
			args: args{
				t: time.Date(2024, 4, 28, 15, 16, 17, 0, time.Local),
			},
			want: "2024-W17-7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ISOInfo(tt.args.t); got != tt.want {
				t.Errorf("ISOInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeOfISOInfo(t *testing.T) {
	type args struct {
		info string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "2024-W17-7, 2024-04-28",
			args: args{
				info: "2024-W17-7",
			},
			want:    time.Date(2024, 4, 28, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name: "2024-W05-2, 2024-01-30",
			args: args{
				info: "2024-W05-2",
			},
			want:    time.Date(2024, 1, 30, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name: "2024-W5-2, 2024-01-30",
			args: args{
				info: "2024-W5-2",
			},
			want:    time.Date(2024, 1, 30, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name: "2024-W05-1, 2024-01-29",
			args: args{
				info: "2024-W05-1",
			},
			want:    time.Date(2024, 1, 29, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name: "2024-W05-7, 2024-02-04",
			args: args{
				info: "2024-W05-7",
			},
			want:    time.Date(2024, 2, 4, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name: "2024-W01-1, 2024-01-01",
			args: args{
				info: "2024-W01-1",
			},
			want:    time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name: "2003-W52-7, 2003-12-28",
			args: args{
				info: "2003-W52-7",
			},
			want:    time.Date(2003, 12, 28, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name: "2004-W01-1, 2003-12-29",
			args: args{
				info: "2004-W01-1",
			},
			want:    time.Date(2003, 12, 29, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name: "2004-W53-7, 2005-01-02",
			args: args{
				info: "2004-W53-7",
			},
			want:    time.Date(2005, 1, 2, 0, 0, 0, 0, time.Local),
			wantErr: false,
		},
		{
			name: "2100-W01-1, error",
			args: args{
				info: "2100-W01-1",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "1999-W01-1, error",
			args: args{
				info: "1999-W01-1",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "2004-W54-1, error",
			args: args{
				info: "2004-W54-1",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "2003-W53-1, error",
			args: args{
				info: "2003-W53-1",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "2024-W01-0, error",
			args: args{
				info: "2024-W01-0",
			},
			want:    time.Time{},
			wantErr: true,
		},
		{
			name: "2024-W01-8, error",
			args: args{
				info: "2024-W01-8",
			},
			want:    time.Time{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TimeOfISOInfo(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeOfISOInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TimeOfISOInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestISOEndOfYear(t *testing.T) {
	type args struct {
		year int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "2003, 2003-12-28",
			args: args{
				year: 2003,
			},
			want: time.Date(2003, 12, 28, 0, 0, 0, 0, time.Local),
		},
		{
			name: "2004, 2005-01-02",
			args: args{
				year: 2004,
			},
			want: time.Date(2005, 1, 2, 0, 0, 0, 0, time.Local),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ISOEndOfYear(tt.args.year); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ISOEndOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
