package main

import (
	"reflect"
	"testing"
)

func Test_neighbors(t *testing.T) {
	type args struct {
		c              *coordinate
		prev           *coordinate
		allowDiagonals bool
		used           [][]bool
		data           [][]rune
	}
	tests := []struct {
		name    string
		args    args
		want    []coordinate
		wantErr bool
	}{
		{name: "no diagonals",
			args: args{
				c:              &coordinate{1, 1},
				prev:           nil,
				data:           [][]rune{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}},
				allowDiagonals: false,
				used:           nil,
			},
			want:    []coordinate{{0, 1}, {2, 1}, {1, 0}, {1, 2}},
			wantErr: false},
		{name: "no diagonals, prev",
			args: args{
				c:              &coordinate{1, 1},
				prev:           &coordinate{0, 1},
				data:           [][]rune{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}},
				allowDiagonals: false,
				used:           nil,
			},
			want:    []coordinate{{2, 1}, {1, 0}, {1, 2}},
			wantErr: false},
		{name: "diagonals",
			args: args{
				c:              &coordinate{1, 1},
				prev:           nil,
				data:           [][]rune{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}},
				allowDiagonals: true,
				used:           nil,
			},
			want: []coordinate{{0, 1}, {2, 1}, {1, 0}, {1, 2},
				{0, 0}, {2, 0}, {0, 2}, {2, 2}},
			wantErr: false},
		{name: "diagonals used",
			args: args{
				c:              &coordinate{1, 1},
				prev:           nil,
				data:           [][]rune{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}},
				allowDiagonals: true,
				used:           [][]bool{{true, true, true}, {true, false, false}, nil},
			},
			want: []coordinate{{2, 1}, {1, 2},
				{2, 0}, {2, 2}},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := neighbors(tt.args.c, tt.args.prev, tt.args.allowDiagonals, tt.args.used, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("neighbors() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("neighbors() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindWord(t *testing.T) {
	type args struct {
		data [][]rune
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "bucks",
			args:    args{data: [][]rune{{'s', 'k', 'c'}, {'x', 'b', 'u'}, {'y', 'a', 'z'}}},
			want:    "bucks",
			wantErr: false,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindWord(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FindWord() got = %v, want %v", got, tt.want)
			}
		})
	}
}
