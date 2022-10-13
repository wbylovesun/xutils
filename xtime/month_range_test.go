package xtime

import (
	"reflect"
	"testing"
	"time"
)

func TestMonthRange_Slice(t *testing.T) {
	type fields struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []int
	}{
		{
			name: "202201-202206",
			fields: fields{
				from: time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local),
				to:   time.Date(2022, 6, 1, 0, 0, 0, 0, time.Local),
			},
			want: []int{202201, 202202, 202203, 202204, 202205, 202206},
		},
		{
			name: "202007-202206",
			fields: fields{
				from: time.Date(2020, 7, 1, 0, 0, 0, 0, time.Local),
				to:   time.Date(2022, 6, 1, 0, 0, 0, 0, time.Local),
			},
			want: []int{
				202007, 202008, 202009, 202010, 202011, 202012,
				202101, 202102, 202103, 202104, 202105, 202106,
				202107, 202108, 202109, 202110, 202111, 202112,
				202201, 202202, 202203, 202204, 202205, 202206,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &MonthRange{
				from: tt.fields.from,
				to:   tt.fields.to,
			}
			if got := mr.Slice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Slice() = %v, want %v", got, tt.want)
			}
		})
	}
}
