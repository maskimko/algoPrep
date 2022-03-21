package main

import "testing"

func Test_wordLadder(t *testing.T) {
	type args struct {
		begin    string
		end      string
		wordList []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{
				begin:    "cold",
				end:      "warm",
				wordList: []string{"cold", "gold", "cord", "card", "ward", "warm", "tard", "sold"},
			},
			want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordLadder(tt.args.begin, tt.args.end, tt.args.wordList); got != tt.want {
				t.Errorf("wordLadder() = %v, want %v", got, tt.want)
			}
		})
	}
}
