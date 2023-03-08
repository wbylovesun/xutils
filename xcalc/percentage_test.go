package xcalc

import (
	"math"
	"testing"
)

func TestPercentageWithoutDigit(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "int/int",
			args: args{
				numerator:   80,
				denominator: 110,
			},
			want:    72,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PercentageNoDigit(tt.args.numerator, tt.args.denominator)
			if (err != nil) != tt.wantErr {
				t.Errorf("PercentageNoDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PercentageNoDigit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64PercentageWithoutDigit(t *testing.T) {
	type args struct {
		numerator   float64
		denominator float64
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "int/int",
			args: args{
				numerator:   80.8,
				denominator: 110,
			},
			want:    73,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PercentageNoDigit(tt.args.numerator, tt.args.denominator)
			if (err != nil) != tt.wantErr {
				t.Errorf("PercentageNoDigit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PercentageNoDigit() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercentage(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "8/11",
			args: args{
				numerator:   8,
				denominator: 11,
			},
			want:    72.73,
			wantErr: false,
		},
		{
			name: "8/21",
			args: args{
				numerator:   8,
				denominator: 21,
			},
			want:    38.1,
			wantErr: false,
		},
		{
			name: "8/22",
			args: args{
				numerator:   8,
				denominator: 22,
			},
			want:    36.36,
			wantErr: false,
		},
		{
			name: "8/26",
			args: args{
				numerator:   8,
				denominator: 26,
			},
			want:    30.77,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Percentage(tt.args.numerator, tt.args.denominator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Percentage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Percentage() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercentagePrecision(t *testing.T) {
	type args struct {
		numerator   int
		denominator int
		precision   int
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "8/11",
			args: args{
				numerator:   8,
				denominator: 11,
				precision:   4,
			},
			want:    72.7273,
			wantErr: false,
		},
		{
			name: "8/21",
			args: args{
				numerator:   8,
				denominator: 21,
				precision:   4,
			},
			want:    38.0952,
			wantErr: false,
		},
		{
			name: "8/22",
			args: args{
				numerator:   8,
				denominator: 22,
				precision:   7,
			},
			want:    36.3636364,
			wantErr: false,
		},
		{
			name: "8/31",
			args: args{
				numerator:   8,
				denominator: 31,
				precision:   4,
			},
			want:    25.8065,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PercentagePrecision(tt.args.numerator, tt.args.denominator, tt.args.precision)
			if (err != nil) != tt.wantErr {
				t.Errorf("PercentagePrecision() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PercentagePrecision() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPercentageNoDigitWithSuppression(t *testing.T) {
	type args struct {
		numerator   float64
		denominator float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "inf/0",
			args: args{
				numerator:   math.Inf(1),
				denominator: 0,
			},
			want: 0,
		},
		{
			name: "0/inf",
			args: args{
				numerator:   0,
				denominator: math.Inf(1),
			},
			want: 0,
		},
		{
			name: "nan/0",
			args: args{
				numerator:   math.NaN(),
				denominator: 0,
			},
			want: 0,
		},
		{
			name: "0/nan",
			args: args{
				numerator:   0,
				denominator: math.NaN(),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PercentageNoDigitWithSuppression(tt.args.numerator, tt.args.denominator); got != tt.want {
				t.Errorf("PercentageNoDigitWithSuppression() = %v, want %v", got, tt.want)
			}
		})
	}
}
