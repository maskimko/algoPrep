package main

import "testing"

func Test_isValidCourseSchedule(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidCourseSchedule(tt.args.n, tt.args.prerequisites); got != tt.want {
				t.Errorf("isValidCourseSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}
