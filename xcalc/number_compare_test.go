package xcalc

import "testing"

func TestMin(t *testing.T) {
	type args struct {
		z []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "zero item",
			args: args{z: nil},
			want: 0,
		},
		{
			name: "1 item",
			args: args{z: []float64{1.2}},
			want: 1.2,
		},
		{
			name: "2 items",
			args: args{z: []float64{1.2, 1.3}},
			want: 1.2,
		},
		{
			name: "more than 2 items",
			args: args{z: []float64{1.2, 1.3, -8, -20.8}},
			want: -20.8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.z...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		z []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "zero item",
			args: args{z: nil},
			want: 0,
		},
		{
			name: "1 item",
			args: args{z: []float64{1.2}},
			want: 1.2,
		},
		{
			name: "2 items",
			args: args{z: []float64{1.2, 1.3}},
			want: 1.3,
		},
		{
			name: "more than 2 items",
			args: args{z: []float64{1.2, 1.3, -8, -20.8}},
			want: 1.3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.z...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromiseGte0(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "-10.0 output 0",
			args: args{x: -10.0},
			want: 0,
		},
		{
			name: "10.1 output 10.1",
			args: args{x: 10.1},
			want: 10.1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PromiseGte0(tt.args.x); got != tt.want {
				t.Errorf("PromiseGte0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPromiseLte100(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "120.10 output 100",
			args: args{x: 120.10},
			want: 100,
		},
		{
			name: "80.05 output 80.05",
			args: args{x: 80.05},
			want: 80.05,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PromiseLte100(tt.args.x); got != tt.want {
				t.Errorf("PromiseLte100() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsGte(t *testing.T) {
	type args struct {
		x float64
		y float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1 is gte -1",
			args: args{x: 1, y: -1},
			want: true,
		},
		{
			name: "0 is gte0 0",
			args: args{
				x: 0,
				y: 0,
			},
			want: true,
		},
		{
			name: "0 is not gte 1",
			args: args{
				x: 0,
				y: 1,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsGte(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("IsGte() = %v, want %v", got, tt.want)
			}
		})
	}
}
