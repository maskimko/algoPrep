package main

import "testing"

func Test_numSteps(t *testing.T) {
	type args struct {
		combo         string
		trappedCombos []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{
				combo:         "0202",
				trappedCombos: []string{"0201", "0101", "0102", "1212", "2002"},
			},
			want: 6},
		{name: "case 3",
			args: args{
				combo:         "1111",
				trappedCombos: []string{"0111", "2111", "1011", "1211", "1101", "1121", "1110", "1112"},
			},
			want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSteps(tt.args.combo, tt.args.trappedCombos); got != tt.want {
				t.Errorf("numSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
