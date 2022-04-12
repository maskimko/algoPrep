package autoscalePolicyUtilizationCheck

import "testing"

func Test_autoScale(t *testing.T) {
	type args struct {
		utilization []int
		instances   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "from youtube",
			args: args{
				utilization: []int{25, 23, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 76, 80},
				instances:   2,
			},
			want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := autoScale(tt.args.utilization, tt.args.instances); got != tt.want {
				t.Errorf("autoScale() = %v, want %v", got, tt.want)
			}
		})
	}
}
