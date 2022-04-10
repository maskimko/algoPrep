package main

import (
	"strings"
	"testing"
)

func Test_middleOfLinkedList(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{input: "0 1 2 3 4"},
			want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headIdx := 0
			list := buildList(strings.Split(tt.args.input, " "), &headIdx)
			if got := middleOfLinkedList(list); got != tt.want {
				t.Errorf("middleOfLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}
