package xslice

import (
	"reflect"
	"testing"
)

func TestToString(t *testing.T) {
	type args[T interface{ NumberSliceElementType | UnsignedNumberSliceElementType }] struct {
		t   []T
		sep string
	}
	type testCase[T interface{ NumberSliceElementType | UnsignedNumberSliceElementType }] struct {
		name string
		args args[T]
		want string
	}
	tests := []testCase[int]{
		{
			name: "Int",
			args: args[int]{
				t:   []int{1, 2, -1, -3, 0},
				sep: ",",
			},
			want: "1,2,-1,-3,0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Join(tt.args.t, tt.args.sep); got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToIntSlice(t *testing.T) {
	type args[T interface{ NumberSliceElementType | UnsignedNumberSliceElementType }] struct {
		t []T
	}
	type testCase[T interface{ NumberSliceElementType | UnsignedNumberSliceElementType }] struct {
		name string
		args args[T]
		want []int
	}
	tests := []testCase[uint64]{
		{
			name: "Uint",
			args: args[uint64]{
				t: []uint64{1, 2, 3, 4, 5000000000000},
			},
			want: []int{1, 2, 3, 4, 5000000000000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToIntSlice(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToIntSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
