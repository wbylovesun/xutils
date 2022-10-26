package xtime

import (
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
			want:            107,
			expectInitError: false,
		},
		{
			name:            "15mins span",
			fields:          fields{span: 15},
			args:            args{i: 533},
			want:            36,
			expectInitError: false,
		},
		{
			name:            "30mins span",
			fields:          fields{span: 30},
			args:            args{i: 533},
			want:            18,
			expectInitError: false,
		},
		{
			name:            "60mins span",
			fields:          fields{span: 60},
			args:            args{i: 533},
			want:            9,
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
			got, err := ts.OfTime(tt.args.i)
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
