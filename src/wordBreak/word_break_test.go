package main

import "testing"

func Test_wordBreak(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "algomonster",
			args: args{
				s:     "algomonster",
				words: []string{"algo", "monster"},
			},
			want: true},
		{name: "aab",
			args: args{
				s:     "aab",
				words: []string{"a", "c"},
			},
			want: false},
		{name: "big a",
			args: args{
				s:     "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
				words: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			},
			want: false},
		{name: "abcd",
			args: struct {
				s     string
				words []string
			}{s: "abcd", words: []string{"a", "abc", "b", "cd"}},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.words); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
