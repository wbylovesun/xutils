package xslice

import (
	"github.com/wbylovesun/xutils/xstrings"
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "length is 0",
			args: args{data: []int{}},
			want: nil,
		},
		{
			name: "[]int{2,3,5}",
			args: args{data: []int{2, 3, 5}},
			want: func() *int { v := 2; return &v }(),
		},
		{
			name: "[]int{2}",
			args: args{data: []int{2}},
			want: func() *int { v := 2; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Min(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "length is 0",
			args: args{data: []int{}},
			want: nil,
		},
		{
			name: "[]int{2,3,5}",
			args: args{data: []int{2, 3, 5}},
			want: func() *int { v := 5; return &v }(),
		},
		{
			name: "[]int{2}",
			args: args{data: []int{2}},
			want: func() *int { v := 2; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxString(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "length is 0",
			args: args{data: []string{}},
			want: nil,
		},
		{
			name: "[]string{hello, world, what}",
			args: args{data: []string{"hello", "world", "what"}},
			want: func() *string { v := "world"; return &v }(),
		},
		{
			name: "[]string{hello}",
			args: args{data: []string{"hello"}},
			want: func() *string { v := "hello"; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsSlice(t *testing.T) {
	type args struct {
		data []int
		ct   []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "containsTrue",
			args: args{
				data: []int{2, 3, 5, 8, 9, 13, 21, 34, 35, 46, 47, 48, 51, 52, 54, 55, 58, 61, 63, 69},
				ct:   []int{2, 48, 63},
			},
			want: true,
		},
		{
			name: "duplicateContainsTrue",
			args: args{
				data: []int{2, 3, 5, 8, 9, 13, 21, 34, 35, 46, 47, 48, 51, 52, 54, 55, 58, 61, 63, 69},
				ct:   []int{2, 48, 5, 63, 5},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsSlice(tt.args.data, tt.args.ct); got != tt.want {
				t.Errorf("ContainsSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBisectionSearch(t *testing.T) {
	type args struct {
		list   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search 8 in []int{2,3,8,10}",
			args: args{
				list:   []int{2, 3, 8, 10},
				target: 8,
			},
			want: 2,
		},
		{
			name: "not found",
			args: args{
				list:   []int{2, 3, 8, 10},
				target: 4,
			},
			want: -1,
		},
		{
			name: "empty slice",
			args: args{
				list:   nil,
				target: 4,
			},
			want: -1,
		},
		{
			name: "only 1 item slice",
			args: args{
				list:   []int{4},
				target: 4,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BisectionSearch(tt.args.list, tt.args.target); got != tt.want {
				t.Errorf("BisectionSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		data []int
		f    func(v int, k int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "filter value lte 10",
			args: args{
				data: []int{4, 6, 7, 10, 11, 13, 21},
				f: func(v int, k int) bool {
					return v <= 10
				},
			},
			want: []int{11, 13, 21},
		},
		{
			name: "filter value gt 10",
			args: args{
				data: []int{4, 6, 7, 10, 11, 13, 21},
				f: func(v int, k int) bool {
					return v > 10
				},
			},
			want: []int{4, 6, 7, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.data, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterString(t *testing.T) {
	type args struct {
		data []string
		f    func(v string, k int) bool
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "filter value not start with hello",
			args: args{
				data: []string{"hello world", "what's up", "hello kitty", "hello peppa", "susy sheep", "hello baby"},
				f: func(v string, k int) bool {
					return !xstrings.StartsWith(v, "hello")
				},
			},
			want: []string{"hello world", "hello kitty", "hello peppa", "hello baby"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.data, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWalk(t *testing.T) {
	type args struct {
		data []int
		f    func(v int, k int) int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "multiple 2",
			args: args{
				data: []int{2, 3, 8, 12},
				f: func(v int, k int) int {
					return v * 2
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var x []int
			copy(x, tt.args.data)
			Walk(tt.args.data, tt.args.f)
			for k, v := range x {
				if v*2 != tt.args.data[k] {
					t.Errorf("wants %v, got %v", v*2, tt.args.data[k])
				}
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args struct {
		f    func(v int) int
		data [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "multiple 2",
			args: args{
				f: func(v int) int {
					return v * 2
				},
				data: [][]int{{1, 2, 3, 5, 7}, {4}},
			},
			want: []int{2, 4, 6, 10, 14, 8},
		},
		{
			name: "sub 2",
			args: args{
				f: func(v int) int {
					return v - 2
				},
				data: [][]int{{1, 2, 3, 5, 7}, {4}},
			},
			want: []int{-1, 0, 1, 3, 5, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.f, tt.args.data...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
