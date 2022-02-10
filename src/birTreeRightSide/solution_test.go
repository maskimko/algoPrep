package main

import (
	"reflect"
	"strings"
	"testing"
)

func Test_binaryTreeRightSideView(t *testing.T) {
	ex1Idx := 0
	ex1 := buildTree(strings.Split("1 2 4 x 7 x x 5 x x 3 x 6 x x", " "), &ex1Idx)
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "example 1",
			args: args{root: ex1},
			want: []int{1, 3, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreeRightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("binaryTreeRightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
