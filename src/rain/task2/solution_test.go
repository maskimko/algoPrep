package main

import "testing"

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{[]int{7,
				3,
				7,
				3,
				1,
				3,
				4,
				1}},
			want: 5},
		{name: "my test",
			args: args{[]int{1, 2, 1, 3, 1, 1, 2, 1, 4, 3, 3, 3, 2, 1, 2, 2, 2, 4, 1, 2, 3, 4, 3, 3, 1, 2, 4}},
			want: 4},
		{name: "case 2",
			args: args{[]int{2, 1, 1, 3, 2, 1, 1, 3}},
			want: 3},
		{name: "case 3",
			args: args{[]int{7, 5, 2, 7, 2, 7, 4, 7}},
			want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
