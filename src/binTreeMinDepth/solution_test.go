package main

import (
	"strings"
	"testing"
)

func Test_binaryTreeMinDepth(t *testing.T) {
	rIdx := 0
	zero := buildTree(strings.Split("0 x x", " "), &rIdx)
	rIdx = 0
	example := buildTree(strings.Split("1 2 4 x 7 x x 5 x x 3 x 6 x x", " "), &rIdx)
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "zero",
			args: args{root: zero},
			want: 0},
		{name: "example",
			args: args{example},
			want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryTreeMinDepth(tt.args.root); got != tt.want {
				t.Errorf("binaryTreeMinDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
