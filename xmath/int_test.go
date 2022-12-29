package xmath

import "testing"

func TestRound(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "floor-1",
			args: args{x: 0.45},
			want: 0,
		},
		{
			name: "floor-2",
			args: args{x: 0.49},
			want: 0,
		},
		{
			name: "floor-3",
			args: args{x: -1.5},
			want: -2,
		},
		{
			name: "ceil-2",
			args: args{x: -1.4},
			want: -1,
		},
		{
			name: "ceil-1",
			args: args{x: 0.5},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.x); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
