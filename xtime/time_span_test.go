package xtime

import (
	"reflect"
	"testing"
)

func Test_timeSpan_Index(t *testing.T) {
	type fields struct {
		span      int
		timeSpans []int
		timeSegs  []string
	}
	type args struct {
		i int
	}
	tests := []struct {
		name            string
		fields          fields
		args            args
		want            int
		expectInitError bool
	}{
		{
			name: "5mins span",
			fields: fields{
				span:      5,
				timeSpans: nil,
				timeSegs:  nil,
			},
			args:            args{i: 533},
			want:            106,
			expectInitError: false,
		},
		{
			name:            "15mins span",
			fields:          fields{span: 15},
			args:            args{i: 533},
			want:            35,
			expectInitError: false,
		},
		{
			name:            "30mins span",
			fields:          fields{span: 30},
			args:            args{i: 533},
			want:            17,
			expectInitError: false,
		},
		{
			name:            "60mins span",
			fields:          fields{span: 60},
			args:            args{i: 533},
			want:            8,
			expectInitError: false,
		},
		{
			name:            "90mins span",
			fields:          fields{span: 90},
			args:            args{i: 533},
			want:            5,
			expectInitError: false,
		},
		{
			name:            "120mins span",
			fields:          fields{span: 120},
			args:            args{i: 533},
			want:            4,
			expectInitError: false,
		},
		{
			name:            "180mins span",
			fields:          fields{span: 180},
			args:            args{i: 533},
			want:            2,
			expectInitError: false,
		},
		{
			name:            "240mins span",
			fields:          fields{span: 240},
			args:            args{i: 533},
			want:            2,
			expectInitError: false,
		},
		{
			name:            "360mins span",
			fields:          fields{span: 360},
			args:            args{i: 533},
			want:            1,
			expectInitError: false,
		},
		{
			name:            "720mins span",
			fields:          fields{span: 720},
			args:            args{i: 533},
			want:            0,
			expectInitError: false,
		},
		{
			name:            "45mins - invalid span",
			fields:          fields{span: 45},
			args:            args{i: 533},
			want:            0,
			expectInitError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, err := NewTimeSpan(tt.fields.span)
			if tt.expectInitError {
				if err == nil {
					t.Errorf("Index() does not expect error, but got %v", err)
				}
				return
			}
			if got := ts.Index(tt.args.i); got != tt.want {
				t.Errorf("Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeSpan_Of(t *testing.T) {
	type fields struct {
		span int
	}
	type args struct {
		i int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "5mins span",
			fields: fields{span: 5},
			args:   args{i: 533},
			want:   "08:55",
		},
		{
			name:   "15mins span",
			fields: fields{span: 15},
			args:   args{i: 533},
			want:   "09:00",
		},
		{
			name:   "30mins span",
			fields: fields{span: 30},
			args:   args{i: 533},
			want:   "09:00",
		},
		{
			name:   "60mins span",
			fields: fields{span: 60},
			args:   args{i: 533},
			want:   "09:00",
		},
		{
			name:   "90mins span",
			fields: fields{span: 90},
			args:   args{i: 533},
			want:   "09:00",
		},
		{
			name:   "120mins span",
			fields: fields{span: 120},
			args:   args{i: 533},
			want:   "10:00",
		},
		{
			name:   "180mins span",
			fields: fields{span: 180},
			args:   args{i: 533},
			want:   "09:00",
		},
		{
			name:   "240mins span",
			fields: fields{span: 240},
			args:   args{i: 533},
			want:   "12:00",
		},
		{
			name:   "480mins span",
			fields: fields{span: 480},
			args:   args{i: 533},
			want:   "16:00",
		},
		{
			name:   "720mins span",
			fields: fields{span: 720},
			args:   args{i: 533},
			want:   "12:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, _ := NewTimeSpan(tt.fields.span)
			if got := ts.Of(tt.args.i); got != tt.want {
				t.Errorf("Of() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeSpan_OfTime(t *testing.T) {
	type fields struct {
		span      int
		timeSpans []int
		timeSegs  []string
	}
	type args struct {
		i string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "5mins span",
			fields: fields{span: 5},
			args:   args{i: "08:53"},
			want:   "08:55",
		},
		{
			name:   "15mins span",
			fields: fields{span: 15},
			args:   args{i: "08:21"},
			want:   "08:30",
		},
		{
			name:   "30mins span",
			fields: fields{span: 30},
			args:   args{i: "08:00"},
			want:   "08:30",
		},
		{
			name:   "30mins-08:29 span",
			fields: fields{span: 30},
			args:   args{i: "08:29"},
			want:   "08:30",
		},
		{
			name:   "30mins-08:30 span",
			fields: fields{span: 30},
			args:   args{i: "08:30"},
			want:   "09:00",
		},
		{
			name:   "60mins span",
			fields: fields{span: 60},
			args:   args{i: "08:00"},
			want:   "09:00",
		},
		{
			name:   "60mins-08:59 span",
			fields: fields{span: 60},
			args:   args{i: "08:59"},
			want:   "09:00",
		},
		{
			name:   "60mins-09:00 span",
			fields: fields{span: 60},
			args:   args{i: "09:00"},
			want:   "10:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, _ := NewTimeSpan(tt.fields.span)
			got, err := ts.AlignToNextSeg(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("OfTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("OfTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeSpan_isValidHourMinute(t *testing.T) {
	type fields struct {
		span int
	}
	type args struct {
		i string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "1:0",
			fields: fields{span: 5},
			args:   args{i: "1:0"},
			want:   true,
		},
		{
			name:   "1:5",
			fields: fields{span: 5},
			args:   args{i: "1:5"},
			want:   true,
		},
		{
			name:   "1:59",
			fields: fields{span: 5},
			args:   args{i: "1:59"},
			want:   true,
		},
		{
			name:   "1:60",
			fields: fields{span: 5},
			args:   args{i: "1:60"},
			want:   false,
		},
		{
			name:   "01:0",
			fields: fields{span: 5},
			args:   args{i: "01:0"},
			want:   true,
		},
		{
			name:   "01:5",
			fields: fields{span: 5},
			args:   args{i: "01:5"},
			want:   true,
		},
		{
			name:   "01:00",
			fields: fields{span: 5},
			args:   args{i: "01:00"},
			want:   true,
		},
		{
			name:   "01:59",
			fields: fields{span: 5},
			args:   args{i: "01:59"},
			want:   true,
		},
		{
			name:   "01:60",
			fields: fields{span: 5},
			args:   args{i: "01:60"},
			want:   false,
		},
		{
			name:   "a01:60",
			fields: fields{span: 5},
			args:   args{i: "a01:60"},
			want:   false,
		},
		{
			name:   "01:60b",
			fields: fields{span: 5},
			args:   args{i: "01:60b"},
			want:   false,
		},
		{
			name:   "01:6b",
			fields: fields{span: 5},
			args:   args{i: "01:6b"},
			want:   false,
		},
		{
			name:   "a1:6",
			fields: fields{span: 5},
			args:   args{i: "a1:6"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, _ := NewTimeSpan(tt.fields.span)
			if got := ts.isValidHourMinute(tt.args.i); got != tt.want {
				t.Errorf("isValidHourMinute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_timeSpan_Fill(t *testing.T) {
	type fields struct {
		span int
	}
	type args struct {
		m    map[PointOfTimeSlice]int
		pots []PointOfTimeSlice
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "120min",
			fields: fields{
				span: 120,
			},
			args: args{
				m:    map[PointOfTimeSlice]int{"02:00": 10, "06:00": 20, "12:00": 30, "18:00": 60},
				pots: []PointOfTimeSlice{"16:26"},
			},
			want: 10,
		},
		{
			name:   "till per 60min",
			fields: fields{span: 60},
			args: args{
				m:    map[PointOfTimeSlice]int{"04:00": 50, "08:00": 90, "12:00": 70, "15:00": 30},
				pots: []PointOfTimeSlice{"16:26"},
			},
			want: 17,
		},
		{
			name:   "till per 90min",
			fields: fields{span: 90},
			args: args{
				m:    map[PointOfTimeSlice]int{"03:00": 50, "09:00": 90, "13:30": 70, "15:00": 30},
				pots: []PointOfTimeSlice{"16:26"},
			},
			want: 11,
		},
		{
			name:   "fullDay per 60min",
			fields: fields{span: 60},
			args: args{
				m:    map[PointOfTimeSlice]int{"04:00": 50, "08:00": 90, "12:00": 70, "15:00": 30},
				pots: []PointOfTimeSlice{"24:00"},
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts, _ := NewTimeSpan(tt.fields.span)
			if got := ts.Fill(tt.args.m, tt.args.pots...); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("Fill() = %v, want %v, gotValue=%v", len(got), tt.want, got)
			}
		})
	}
}
