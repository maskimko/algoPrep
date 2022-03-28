package main

import (
	"strings"
	"testing"
)

func Test_taskScheduling2(t *testing.T) {
	type args struct {
		tasks        []string
		times        []int
		requirements [][]string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 3",
			args: args{
				tasks: strings.Split("abbreviate bricks cardinals dextrous fibre green height", " "),
				times: []int{1, 1, 1, 1, 1, 100, 1},
				requirements: [][]string{
					{"abbreviate", "bricks"},
					{"cardinals", "bricks"},
					{"dextrous", "bricks"},
					{"bricks", "fibre"},
					{"green", "fibre"},
				},
			},
			want: 101},
		{name: "case 5",
			args: args{
				tasks: strings.Split("aaaabc bbbbbabc cccccabc ddddddabc", " "),
				times: []int{1, 2, 3, 1},
				requirements: [][]string{
					{"aaaabc", "bbbbbabc"},
					{"aaaabc", "cccccabc"},
					{"bbbbbabc", "ddddddabc"},
					{"cccccabc", "ddddddabc"},
				},
			},
			want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := taskScheduling2(tt.args.tasks, tt.args.times, tt.args.requirements); got != tt.want {
				t.Errorf("taskScheduling2() = %v, want %v", got, tt.want)
			}
		})
	}
}
