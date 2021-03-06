package main

import "testing"

func Test_minimumTotal(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1",
			args{[][]int{
				{2},
				{3, 4},
				{5, 6, 7},
				{4, 1, 8, 3}}},
			11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotal(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumTotalBottomUp(t *testing.T) {
	type args struct {
		triangle [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"case 1",
			args{[][]int{
				{2},
				{3, 4},
				{5, 6, 7},
				{4, 1, 8, 3}}},
			11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumTotalBottomUp(tt.args.triangle); got != tt.want {
				t.Errorf("minimumTotalBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
