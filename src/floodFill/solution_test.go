package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_floodFill(t *testing.T) {
	type args struct {
		r           int
		c           int
		replacement int
		image       [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "case 1",
			args: args{
				r:           2,
				c:           2,
				replacement: 9,
				image: [][]int{
					{0, 1, 3, 4, 1},
					{3, 8, 8, 3, 3},
					{6, 7, 8, 8, 3},
					{12, 2, 8, 9, 1},
					{12, 3, 1, 3, 2},
				},
			},
			want: [][]int{
				{0, 1, 3, 4, 1},
				{3, 9, 9, 3, 3},
				{6, 7, 9, 9, 3},
				{12, 2, 9, 9, 1},
				{12, 3, 1, 3, 2},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floodFill(tt.args.r, tt.args.c, tt.args.replacement, tt.args.image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_str_floodFill(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "case 1",
			args: args{
				input: `2
2
9
5
0 1 3 4 1
3 8 8 3 3
6 7 8 8 3
12 2 8 9 1
12 3 1 3 2
`,
			}, want: [][]int{
				{0, 1, 3, 4, 1},
				{3, 9, 9, 3, 3},
				{6, 7, 9, 9, 3},
				{12, 2, 9, 9, 1},
				{12, 3, 1, 3, 2},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, c, replacement, image := processInput(strings.NewReader(tt.args.input))
			if got := floodFill(r, c, replacement, image); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
