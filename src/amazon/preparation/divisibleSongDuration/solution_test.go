package divisibleSongDuration

import (
	"testing"
)

func Test_numPairsDivisibleBy60(t *testing.T) {
	type args struct {
		time []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test 1",
			args: args{[]int{30, 20, 150, 100, 40}},
			want: 3},
		{name: "test 2",
			args: args{[]int{60, 60, 60}},
			want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPairsDivisibleBy60(tt.args.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkPermutations(t *testing.T) {
	type args struct {
		time []int
		j    int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "different 1",
			args: args{time: []int{1, 2, 3, 20, 20, 20, 25, 26, 27, 40, 40, 40, 41, 42, 43},
				j: 3,
				k: 11},
			want: 9},
		{name: "different 2",
			args: args{time: []int{1, 2, 3, 20, 20, 25, 26, 27, 40, 40, 40, 40, 41, 42, 43},
				j: 3,
				k: 11},
			want: 8},
		{name: "same 1",
			args: args{time: []int{1, 2, 3, 30, 30, 30, 41, 42, 43},
				j: 3,
				k: 5},
			want: 3},
		{name: "same 2",
			args: args{time: []int{1, 2, 3, 30, 30, 30, 30, 41, 42, 43},
				j: 3,
				k: 6},
			want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPermutations(tt.args.time, tt.args.j, tt.args.k); got != tt.want {
				t.Errorf("checkPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}
