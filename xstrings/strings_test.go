package xstrings

import "testing"

func TestLcfirst(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Hello",
			args: args{s: "Hello"},
			want: "hello",
		},
		{
			name: "我不会转换",
			args: args{s: "我 never be lower"},
			want: "我 never be lower",
		},
		{
			name: "never be lower",
			args: args{s: "never be lower"},
			want: "never be lower",
		},
		{
			name: "blank string",
			args: args{s: ""},
			want: "",
		},
		{
			name: "one upper case char",
			args: args{s: "H"},
			want: "h",
		},
		{
			name: "chinese sentences start with upper case char",
			args: args{s: "En,我知道了"},
			want: "en,我知道了",
		},
		{
			name: "mixed characters start with upper case char",
			args: args{s: "E我8"},
			want: "e我8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lcfirst(tt.args.s); got != tt.want {
				t.Errorf("Lcfirst() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartsWith(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "original string is empty",
			args: args{s: "", p: "xx"},
			want: false,
		},
		{
			name: "match string is empty",
			args: args{
				s: "xx",
				p: "",
			},
			want: false,
		},
		{
			name: "original string length is lt match string",
			args: args{
				s: "hello",
				p: "helloworld",
			},
			want: false,
		},
		{
			name: "does not match",
			args: args{
				s: "x11",
				p: "x21",
			},
			want: false,
		},
		{
			name: "match",
			args: args{
				s: "helloworld",
				p: "hello",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartsWith(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "original string length is less than match string",
			args: args{
				s: "hello",
				p: "hello world",
			},
			want: false,
		},
		{
			name: "match",
			args: args{
				s: "This sentence ends with hello",
				p: "hello",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndsWith(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("EndsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}
