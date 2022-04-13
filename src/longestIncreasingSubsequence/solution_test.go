package main

import "testing"

func Test_longestSubLen(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 3",
			args: args{[]int{10, 22, 9, 33, 21, 50, 41, 60, 80}},
			want: 6},
		{name: "case 4",
			args: args{[]int{5, 46, 85, 26, 1, 5, 78, 45, 122, 56, 47, 26}},
			want: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubLen(tt.args.nums); got != tt.want {
				t.Errorf("longestSubLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
