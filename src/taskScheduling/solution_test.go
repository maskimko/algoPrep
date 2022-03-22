package main

import (
	"reflect"
	"testing"
)

func Test_taskScheduling(t *testing.T) {
	type args struct {
		tasks        []string
		requirements [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "case 1",
			args: args{
				tasks: []string{"a", "b", "c", "d"},
				requirements: [][]string{
					{"a", "b"},
					{"c", "b"},
					{"b", "d"},
				},
			},
			want: []string{"a", "c", "b", "d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := taskScheduling(tt.args.tasks, tt.args.requirements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taskScheduling() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taskSchedulingDFS(t *testing.T) {
	type args struct {
		tasks        []string
		requirements [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "case 1",
			args: args{
				tasks: []string{"a", "b", "c", "d"},
				requirements: [][]string{
					{"a", "b"},
					{"c", "b"},
					{"b", "d"},
				},
			},
			want: []string{"a", "c", "b", "d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := taskSchedulingDFS(tt.args.tasks, tt.args.requirements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taskSchedulingDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
