package xcalc

import "testing"

func TestFormatAmountToYuan(t *testing.T) {
	type args[T AmountType] struct {
		amount T
	}
	tests := []struct {
		name string
		args args[float64]
		want float64
	}{
		{
			name: "float64",
			args: args[float64]{amount: 2.1999},
			want: 2.20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatAmountToYuan(tt.args.amount); got != tt.want {
				t.Errorf("FormatAmountToYuan() = %v, want %v", got, tt.want)
			}
		})
	}
}
