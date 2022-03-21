package main

import "testing"

func Test_minimumChunksRequired(t *testing.T) {
	type args struct {
		totalPackets   int64
		uploadedChunks [][]int64
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "case 1",
			args: struct {
				totalPackets   int64
				uploadedChunks [][]int64
			}{totalPackets: 18, uploadedChunks: [][]int64{
				{9, 17},
			}}},
		{name: "case 2",
			args: struct {
				totalPackets   int64
				uploadedChunks [][]int64
			}{totalPackets: 10, uploadedChunks: [][]int64{
				{1, 2},
				{9, 10},
			}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumChunksRequired(tt.args.totalPackets, tt.args.uploadedChunks); got != tt.want {
				t.Errorf("minimumChunksRequired() = %v, want %v", got, tt.want)
			}
		})
	}
}
