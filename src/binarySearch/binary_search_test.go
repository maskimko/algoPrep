package main

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		arr    []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "non even",
			args: args{
				arr:    []int{1, 2, 3, 4, 5, 6, 7},
				target: 5,
			},
			want: 4,
		},
		{name: "even",
			args: args{
				arr:    []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
				target: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.arr, tt.args.target); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
