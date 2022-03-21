package main

import "testing"

func Test_numSteps(t *testing.T) {
	type args struct {
		initPos [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{initPos: [][]int{
				{4, 1, 3},
				{2, 0, 5},
			}},
			want: 5},
		{name: "case 1",
			args: args{initPos: [][]int{
				{1, 2, 3},
				{5, 4, 0},
			}},
			want: -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSteps(tt.args.initPos); got != tt.want {
				t.Errorf("numSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
