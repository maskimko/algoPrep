package main

import (
	"reflect"
	"testing"
)

func Test_moveZeros(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case 2",
			args: args{nums: []int{3, 1, 0, 1, 3, 8, 9}},
			want: []int{3, 1, 1, 3, 8, 9, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeros(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("wrong order %v wanted %v", tt.args.nums, tt.want)
			}
		})
	}
}
