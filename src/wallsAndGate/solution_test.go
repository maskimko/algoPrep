package main

import (
	"bufio"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test_str_mapGateDistances(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "case 1",
			args: args{input: `4
2147483647 -1 0 2147483647
2147483647 2147483647 2147483647 -1
2147483647 -1 2147483647 -1
0 -1 2147483647 2147483647
`},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.args.input))
			scanner.Scan()
			dungeonMapLength, _ := strconv.Atoi(scanner.Text())
			dungeonMap := [][]int{}
			for i := 0; i < dungeonMapLength; i++ {
				scanner.Scan()
				dungeonMap = append(dungeonMap, arrayAtoi(splitWords(scanner.Text())))
			}
			if got := mapGateDistances(dungeonMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapGateDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
