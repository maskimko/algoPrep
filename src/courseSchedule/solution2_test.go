package main

import "testing"

func Test_isValidCourseSchedule2(t *testing.T) {
	type args struct {
		n             int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "case 1",
			args: args{
				n:             2,
				prerequisites: [][]int{{0, 1}},
			},
			want: true},
		{name: "case 2",
			args: args{
				n:             2,
				prerequisites: [][]int{{0, 1}, {1, 0}},
			},
			want: false},
		{name: "case 3",
			args: args{
				n:             4,
				prerequisites: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 1}},
			},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidCourseSchedule2(tt.args.n, tt.args.prerequisites); got != tt.want {
				t.Errorf("isValidCourseSchedule2() = %v, want %v", got, tt.want)
			}
		})
	}
}
