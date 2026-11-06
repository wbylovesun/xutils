package xslice

import (
	"reflect"
	"testing"

	"github.com/wbylovesun/xutils/types"
)

type testUserID int64

func TestToString(t *testing.T) {
	type args[T types.Integer] struct {
		t   []T
		sep string
	}
	type testCase[T types.Integer] struct {
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
	type args[T types.Integer] struct {
		t []T
	}
	type testCase[T types.Integer] struct {
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

func TestJoin(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		got := Join([]float64{3, 3.14, -2.5, 0}, ",")
		if got != "3,3.14,-2.5,0" {
			t.Errorf("Join(float64) = %q, want %q", got, "3,3.14,-2.5,0")
		}
	})

	t.Run("float32", func(t *testing.T) {
		// 可精确表示的 float32 输出正常
		if got := Join([]float32{3.5, -2.5}, ","); got != "3.5,-2.5" {
			t.Errorf("Join(float32) = %q, want %q", got, "3.5,-2.5")
		}
		// float32 按自身 32 位精度输出，不产生 64 位伪影
		if got := Join([]float32{0.1, 3.14}, ","); got != "0.1,3.14" {
			t.Errorf("Join(float32) = %q, want %q", got, "0.1,3.14")
		}
	})

	t.Run("uint", func(t *testing.T) {
		got := Join([]uint64{1, 5000000000000}, ",")
		if got != "1,5000000000000" {
			t.Errorf("Join(uint64) = %q, want %q", got, "1,5000000000000")
		}
	})

	t.Run("int8", func(t *testing.T) {
		got := Join([]int8{-128, 127}, ",")
		if got != "-128,127" {
			t.Errorf("Join(int8) = %q, want %q", got, "-128,127")
		}
	})

	t.Run("单元素无分隔符", func(t *testing.T) {
		if got := Join([]int{42}, ","); got != "42" {
			t.Errorf("Join = %q, want %q", got, "42")
		}
	})

	t.Run("空分隔符", func(t *testing.T) {
		if got := Join([]int{1, 2, 3}, ""); got != "123" {
			t.Errorf("Join = %q, want %q", got, "123")
		}
	})

	t.Run("空切片返回空串", func(t *testing.T) {
		if got := Join([]int{}, ","); got != "" {
			t.Errorf("Join = %q, want 空串", got)
		}
	})

	t.Run("nil返回空串", func(t *testing.T) {
		if got := Join([]int(nil), ","); got != "" {
			t.Errorf("Join = %q, want 空串", got)
		}
	})

	t.Run("具名类型", func(t *testing.T) {
		// Number 带 ~ 后，底层类型为数值的自定义类型可直接使用
		if got := Join([]testUserID{1, 2}, ","); got != "1,2" {
			t.Errorf("Join(testUserID) = %q, want %q", got, "1,2")
		}
	})
}

func TestToNumberSlice(t *testing.T) {
	t.Run("int32到int64", func(t *testing.T) {
		got := ToNumberSlice([]int32{-1, 0, 2}, int64(0))
		if !reflect.DeepEqual(got, []int64{-1, 0, 2}) {
			t.Errorf("ToNumberSlice = %v, want %v", got, []int64{-1, 0, 2})
		}
	})

	t.Run("uint8到int32", func(t *testing.T) {
		got := ToNumberSlice([]uint8{200, 255}, int32(0))
		if !reflect.DeepEqual(got, []int32{200, 255}) {
			t.Errorf("ToNumberSlice = %v, want %v", got, []int32{200, 255})
		}
	})

	// 数值转换遵循 Go 语义：同尺寸整数转换保留位模式，溢出回绕
	t.Run("int32到int8回绕", func(t *testing.T) {
		got := ToNumberSlice([]int32{300}, int8(0))
		if !reflect.DeepEqual(got, []int8{44}) {
			t.Errorf("ToNumberSlice = %v, want %v（300 回绕为 44）", got, []int8{44})
		}
	})

	t.Run("uint32到int32回绕", func(t *testing.T) {
		got := ToNumberSlice([]uint32{4294967295}, int32(0))
		if !reflect.DeepEqual(got, []int32{-1}) {
			t.Errorf("ToNumberSlice = %v, want %v（2^32-1 位模式回绕为 -1）", got, []int32{-1})
		}
	})

	t.Run("int64到int8回绕", func(t *testing.T) {
		got := ToNumberSlice([]int64{130}, int8(0))
		if !reflect.DeepEqual(got, []int8{-126}) {
			t.Errorf("ToNumberSlice = %v, want %v（130 回绕为 -126）", got, []int8{-126})
		}
	})

	t.Run("空切片返回nil", func(t *testing.T) {
		if got := ToNumberSlice([]int32{}, int64(0)); got != nil {
			t.Errorf("ToNumberSlice = %v, want nil", got)
		}
	})

	t.Run("nil返回nil", func(t *testing.T) {
		if got := ToNumberSlice[int32, int64](nil, 0); got != nil {
			t.Errorf("ToNumberSlice = %v, want nil", got)
		}
	})
}

func TestToInt64Slice(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		got := ToInt64Slice([]int{-1, 0, 42})
		if !reflect.DeepEqual(got, []int64{-1, 0, 42}) {
			t.Errorf("ToInt64Slice = %v, want %v", got, []int64{-1, 0, 42})
		}
	})

	t.Run("uint32不回绕", func(t *testing.T) {
		got := ToInt64Slice([]uint32{0, 4294967295})
		if !reflect.DeepEqual(got, []int64{0, 4294967295}) {
			t.Errorf("ToInt64Slice = %v, want %v", got, []int64{0, 4294967295})
		}
	})

	// uint64 最大值超出 int64 表示范围，按位模式回绕为 -1
	t.Run("uint64最大值回绕", func(t *testing.T) {
		got := ToInt64Slice([]uint64{18446744073709551615})
		if !reflect.DeepEqual(got, []int64{-1}) {
			t.Errorf("ToInt64Slice = %v, want %v", got, []int64{-1})
		}
	})

	t.Run("空切片与nil返回nil", func(t *testing.T) {
		if got := ToInt64Slice([]int{}); got != nil {
			t.Errorf("ToInt64Slice([]) = %v, want nil", got)
		}
		if got := ToInt64Slice[int](nil); got != nil {
			t.Errorf("ToInt64Slice(nil) = %v, want nil", got)
		}
	})

	t.Run("具名类型", func(t *testing.T) {
		got := ToInt64Slice([]testUserID{-5, 8})
		if !reflect.DeepEqual(got, []int64{-5, 8}) {
			t.Errorf("ToInt64Slice = %v, want %v", got, []int64{-5, 8})
		}
	})
}

func TestToInt32Slice(t *testing.T) {
	t.Run("int16", func(t *testing.T) {
		got := ToInt32Slice([]int16{-32768, 32767})
		if !reflect.DeepEqual(got, []int32{-32768, 32767}) {
			t.Errorf("ToInt32Slice = %v, want %v", got, []int32{-32768, 32767})
		}
	})

	// 2^31 与 2^32+5 超出 int32 范围，按位模式回绕
	t.Run("uint64溢出回绕", func(t *testing.T) {
		got := ToInt32Slice([]uint64{2147483648, 4294967301})
		if !reflect.DeepEqual(got, []int32{-2147483648, 5}) {
			t.Errorf("ToInt32Slice = %v, want %v", got, []int32{-2147483648, 5})
		}
	})
}

func TestToInt16Slice(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		got := ToInt16Slice([]uint8{0, 255})
		if !reflect.DeepEqual(got, []int16{0, 255}) {
			t.Errorf("ToInt16Slice = %v, want %v", got, []int16{0, 255})
		}
	})

	t.Run("int32溢出回绕", func(t *testing.T) {
		got := ToInt16Slice([]int32{32768, -32769})
		if !reflect.DeepEqual(got, []int16{-32768, 32767}) {
			t.Errorf("ToInt16Slice = %v, want %v（32768→-32768，-32769→32767）", got, []int16{-32768, 32767})
		}
	})
}

func TestToInt8Slice(t *testing.T) {
	t.Run("uint8", func(t *testing.T) {
		got := ToInt8Slice([]uint8{0, 255})
		if !reflect.DeepEqual(got, []int8{0, -1}) {
			t.Errorf("ToInt8Slice = %v, want %v", got, []int8{0, -1})
		}
	})

	t.Run("int64溢出回绕", func(t *testing.T) {
		got := ToInt8Slice([]int64{300, -200})
		if !reflect.DeepEqual(got, []int8{44, 56}) {
			t.Errorf("ToInt8Slice = %v, want %v（300→44，-200→56）", got, []int8{44, 56})
		}
	})

	t.Run("空切片与nil返回nil", func(t *testing.T) {
		if got := ToInt8Slice([]int64{}); got != nil {
			t.Errorf("ToInt8Slice([]) = %v, want nil", got)
		}
		if got := ToInt8Slice[int64](nil); got != nil {
			t.Errorf("ToInt8Slice(nil) = %v, want nil", got)
		}
	})
}
