package main

import "testing"

func Test_decodeWays(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "18",
			args: args{digits: "18"},
			want: 2},
		{name: "123",
			args: args{digits: "123"},
			want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeWays(tt.args.digits); got != tt.want {
				t.Errorf("decodeWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
