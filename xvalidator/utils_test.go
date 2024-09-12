package xvalidator

import (
	"github.com/wbylovesun/xutils/xtime"
	"reflect"
	"testing"
	"time"
)

func Test_isEmbedded(t *testing.T) {
	type args struct {
		struct1 interface{}
		struct2 interface{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				struct1: time.Time{},
				struct2: xtime.JsonShortDate{Time: time.Now()},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEmbedded(reflect.TypeOf(tt.args.struct1), reflect.TypeOf(tt.args.struct2)); got != tt.want {
				t.Errorf("isEmbedded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEmbeddedStruct(t *testing.T) {
	now := time.Now()
	type args struct {
		b            interface{}
		embeddedType interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 bool
	}{
		{
			name: "test1",
			args: args{
				b:            xtime.JsonShortDate{Time: now},
				embeddedType: time.Time{},
			},
			want:  now,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getEmbeddedStruct(tt.args.b, tt.args.embeddedType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEmbeddedStruct() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getEmbeddedStruct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
