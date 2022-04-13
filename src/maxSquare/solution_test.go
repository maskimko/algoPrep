package main

import (
	"strings"
	"testing"
)

func Test_maximalSquare(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{[][]int{
				{1, 0, 1, 0, 0},
				{1, 0, 1, 1, 1},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 1},
			}},
			want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalSquare(tt.args.matrix); got != tt.want {
				t.Errorf("maximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readAndRun(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{input: `4
1 0 1 0 0
1 0 1 1 1
1 1 1 1 0
1 0 0 0 1`},
			want: 4},
		{name: "case 5",
			args: args{input: `6
1 1 1 1 1 0
1 1 1 1 0 0
1 1 1 1 0 0
0 1 1 1 1 0
0 0 0 1 0 0
1 1 1 0 0 0
`},
			want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.args.input)
			if got := readAndRun(reader); got != tt.want {
				t.Errorf("readAndRun() = %v, want %v", got, tt.want)
			}
		})
	}
}
