package main

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "example",
			args: args{s: "aab"},
			//want: [][]string{{"aa","b"},{"a","a","b"}},
			want: [][]string{{"a", "a", "b"}, {"aa", "b"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "single",
			args: args{p: "a"},
			want: true},
		{name: "empty",
			args: args{p: ""},
			want: false},
		{name: "double",
			args: args{p: "bb"},
			want: true},
		{name: "false",
			args: args{p: "ab"},
			want: false},
		{name: "triple",
			args: args{p: "aba"},
			want: true},
		{name: "long not even",
			args: args{p: "qwertyuiopoiuytrewq"},
			want: true},
		{name: "long even",
			args: args{p: "qwertyuiooiuytrewq"},
			want: true},
		{name: "long false",
			args: args{p: "qwertyuiopvoiuytrewq"},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.p); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
