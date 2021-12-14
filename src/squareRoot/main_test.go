package main

import "testing"

func Test_squareRoot(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "16",
			args: args{n: 16},
		want: 4},
		{name: "81",
			args: args{n: 81},
			want: 9},
		{name: "4",
			args: args{n: 4},
			want: 2},
		{name: "256",
			args: args{n: 256},
			want: 16},
		{name: "8",
			args: args{8},
		want: 2},
		{name: "0",
			args: args{n: 0},
			want: 0},
			{name: "1",
				args: args{1},
			want: 1},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := squareRoot(tt.args.n); got != tt.want {
				t.Errorf("squareRoot() = %v, want %v", got, tt.want)
			}
		})
	}
}
