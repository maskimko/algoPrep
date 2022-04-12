package main

import (
	"reflect"
	"testing"
)

func Test_heapTop3(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case 1",
			args: args{arr: []int{3, 1, 2, 10, 33, 100, 20}},
			want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapTop3(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapTop3() = %v, want %v", got, tt.want)
			}
		})
	}
}
