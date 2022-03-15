package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func Test_str_countNumberOfIslands(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{
				input: `6
1 1 1 0 0 0
1 1 1 1 0 0
1 1 1 0 0 0
0 1 0 0 0 0
0 0 0 0 1 0
0 0 0 0 0 0
`,
			},
			want: 2},
		{name: "case 2",
			args: args{input: `2
1 0 1
0 1 0
`},
			want: 3},
		{name: "case 3",
			args: args{input: `3
0 0 0
0 0 0
0 0 0
`},
			want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.args.input))
			scanner.Scan()
			gridLength, _ := strconv.Atoi(scanner.Text())
			grid := [][]int{}
			for i := 0; i < gridLength; i++ {
				scanner.Scan()
				grid = append(grid, arrayAtoi(splitWords(scanner.Text())))
			}
			if got := countNumberOfIslands(grid); got != tt.want {
				t.Errorf("countNumberOfIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
