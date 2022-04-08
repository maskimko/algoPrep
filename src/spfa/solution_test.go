package main

import "testing"

func Test_shortestPathStrInput(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "case 1",
			args: args{input: `4
1 1 2 1
0 1 2 1 3 1
0 1 1 1
1 1
0
3`},
			want:    2,
			wantErr: false},
		{name: "case 7",
			args: args{input: `5
1 100 2 10 3 1
0 100 2 50 3 30 4 9
0 10 1 50 3 10 4 5
0 1 1 30 2 10 4 1
1 9 2 5 3 1
3
2`},
			want:    6,
			wantErr: false},
		{name: "case 2",
			args: args{input: `4
1 1 2 2
0 1 2 0 3 3
0 2 1 0 3 1
1 3 2 1
0
3`},
			want:    2,
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shortestPathStrInput(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortestPathStrInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("shortestPathStrInput() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortestPath(t *testing.T) {
	type args struct {
		graph  map[int][]edge
		root   int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case 1",
			args: args{
				graph: map[int][]edge{
					0: {{1, 1}, {2, 1}},
					1: {{0, 1}, {2, 1}, {3, 1}},
					2: {{0, 1}, {1, 1}},
					3: {{1, 1}}},
				root:   0,
				target: 3,
			},
			want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPath(tt.args.graph, tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("shortestPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
