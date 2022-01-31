package main

import (
	"reflect"
	"testing"
)

func Test_levelOrderTraversal(t *testing.T) {
	i := 0
	root := buildTree(splitWords("1 2 4 x 7 x x 5 x x 3 x 6 x x"), &i)

	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example",
			args: args{root: root},
			want: [][]int{{1}, {2, 3}, {4, 5, 6}, {7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
