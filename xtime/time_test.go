package xtime

import "testing"

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
