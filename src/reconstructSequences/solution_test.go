package main

import "testing"

func Test_sequenceReconstruction(t *testing.T) {
	type args struct {
		original []int
		seqs     [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case 2",
			args: args{
				original: []int{1, 2, 3},
				seqs: [][]int{
					{1, 2},
					{1, 3},
					{2, 3},
				},
			},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sequenceReconstruction(tt.args.original, tt.args.seqs); got != tt.want {
				t.Errorf("sequenceReconstruction() = %v, want %v", got, tt.want)
			}
		})
	}
}
