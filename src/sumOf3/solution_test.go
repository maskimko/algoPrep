package main

import (
	"reflect"
	"testing"
)

func Test_tripletsWithSum0(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "from example 1",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "from example 2",
			args: args{nums: []int{1, -1, 2, -2, 3, -3, 4, -4}},
			want: [][]int{{-4, 1, 3}, {-3, -1, 4}, {-3, 1, 2}, {-2, -1, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tripletsWithSum0(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tripletsWithSum0() = %v, want %v", got, tt.want)
			}
		})
	}
}
