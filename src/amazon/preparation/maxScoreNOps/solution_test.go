package maxScoreNOps

import "testing"

func Test_gcd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{
				a: 15,
				b: 10,
			},
			want: 5},
		{name: "case 2",
			args: args{
				a: 252,
				b: 105,
			},
			want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gcd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxScore(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{nums: []int{1, 2}},
			want: 1},
		{name: "case 2",
			args: args{nums: []int{3, 4, 6, 8}},
			want: 11},
		{name: "test 1",
			args: args{nums: []int{1, 2, 3, 4, 5, 6}},
			want: 14},
		{name: "test 2",
			args: args{nums: []int{537396, 848445, 135528, 543874, 815161, 588627, 215467, 446789}},
			want: 256},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxScore(tt.args.nums); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
