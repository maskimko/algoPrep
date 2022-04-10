package main

import (
	"reflect"
	"testing"
)

func Test_twoSumSorted(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case 1",
			args: args{
				arr:    []int{2, 3, 5, 8, 11, 15},
				target: 5,
			},
			want: []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumSorted(tt.args.arr, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
