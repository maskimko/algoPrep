package main

import (
	"container/heap"
	"reflect"
	"strings"
	"testing"
)

func Test_alienOrder(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case 1",
			args: args{words: []string{"wrt", "wrf", "er", "ett", "rftt"}},
			want: "wertf"},
		{name: "case 3",
			args: args{words: []string{"she", "sell", "seashell", "seashore", "seahorse", "on", "a"}},
			want: "lnrsheoa"},
		{name: "case 5",
			args: args{words: strings.Split("da la na fa fei jia ha hai hang hua ta sha shi si ba", " ")},
			want: ""},
		{name: "simple 1 char",
			args: args{words: []string{"e", "d", "c", "b", "a"}},
			want: "edcba"},
		{name: "simple 2 char",
			args: args{words: []string{"ed", "ec", "db", "da", "ce", "cb", "ca", "ac", "ab"}},
			want: "edcba"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alienOrder(tt.args.words); got != tt.want {
				t.Errorf("alienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pQueue_Pop(t *testing.T) {
	tests := []struct {
		name string
		p    pQueue
		want []rune
	}{
		{name: "PQ test int",
			p:    pQueue{3, 2, 5},
			want: []rune{rune(1), rune(2), rune(3), rune(4), rune(5)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			heap.Init(&tt.p)
			heap.Push(&tt.p, rune(1))
			heap.Push(&tt.p, rune(4))
			got := make([]rune, 0)
			for tt.p.Len() > 0 {
				got = append(got, heap.Pop(&tt.p).(rune))
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pq = %v, want %v", got, tt.want)
			}
		})
	}
}
