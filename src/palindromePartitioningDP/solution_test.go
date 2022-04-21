package main

import "testing"

func Test_partition2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 3",
			args: args{s: "dfxferlmjz"},
			want: 2},
		{name: "case 1",
			args: args{s: "aab"},
			want: 2},
		{name: "my case 1",
			args: args{s: "eabcdedcbaf"},
			want: 2},
		{name: "my case 2",
			args: args{s: "eababaf"},
			want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition2(tt.args.s); got != tt.want {
				t.Errorf("partition2() = %v, want %v", got, tt.want)
			}
		})
	}
}
