package main

import (
	"testing"
)

func Test_newspapersSplit(t *testing.T) {
	type args struct {
		newspapers []int
		coworkers  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				newspapers: []int{7, 2, 5, 10, 8},
				coworkers:  2,
			},
			want: 18,
		},
		{
			name: "example 2",
			args: args{
				newspapers: []int{2, 3, 5, 7},
				coworkers:  3,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newspapersSplit(tt.args.newspapers, tt.args.coworkers); got != tt.want {
				t.Errorf("newspapersSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_assignNewspapers(t *testing.T) {
	type args struct {
		newspapers []int
		capacity   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "assign newspapers",
			args: args{
				newspapers: []int{1, 2, 3, 4, 5, 6},
				capacity:   6,
			},
			want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := assignNewspapers(tt.args.newspapers, tt.args.capacity); got != tt.want {
				t.Errorf("assignNewspapers() = %v, want %v", got, tt.want)
			}
		})
	}
}
