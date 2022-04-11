package main

import (
	"testing"
)

func Test_isAlphaNumeric(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "small",
			args: args{'m'},
			want: true},
		{name: "capital",
			args: args{'M'},
			want: true},
		{name: "number",
			args: args{'1'},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlphaNumeric(tt.args.r); got != tt.want {
				t.Errorf("isAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case 1",
			args: args{s: "Do geese see God?"},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
