package main

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	type args struct {
		points []point
		k      int
	}
	tests := []struct {
		name string
		args args
		want []point
	}{
		{name: "case 1",
			args: args{
				points: []point{{1, 1}, {2, 2}, {3, 3}},
				k:      1,
			},
			want: []point{{1, 1}}},
		{name: "case 2",
			args: args{
				points: []point{{1, 3}, {-2, 2}},
				k:      1,
			},
			want: []point{{-2, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kClosest(tt.args.points, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
