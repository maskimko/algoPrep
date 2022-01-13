package arraySubsets

import (
	"reflect"
	"testing"
)

func Test_getArraySubsets(t *testing.T) {
	type args struct {
		in0 []interface{}
	}
	tests := []struct {
		name string
		args args
		want [][][]interface{}
	}{
		{name: "two",
			args: args{[]interface{}{1, 2}},
			want: [][][]interface{}{{{1, 2}}, {{1}, {2}}},
		},
		{name: "three",
			args: args{[]interface{}{1, 2, 3}},
			want: [][][]interface{}{{{1, 2, 3}}, {{1}, {2, 3}}, {{1}, {2}, {3}}, {{1, 2}, {3}}},
		},
		{name: "five",
			args: args{[]interface{}{1, 2, 3, 4, 5}},
			want: [][][]interface{}{{{1, 2, 3, 4, 5}}, {{1}, {2, 3, 4, 5}}, {{1}, {2}, {3, 4, 5}},
				{{1}, {2}, {3}, {4, 5}}, {{1}, {2}, {3}, {4}, {5}}, {{1}, {2}, {3, 4}, {5}}, {{1}, {2, 3}, {4, 5}},
				{{1}, {2, 3}, {4}, {5}}, {{1}, {2, 3, 4}, {5}}, {{1, 2}, {3, 4, 5}}, {{1, 2}, {3}, {4, 5}},
				{{1, 2}, {3}, {4}, {5}}, {{1, 2}, {3, 4}, {5}}, {{1, 2, 3}, {4, 5}}, {{1, 2, 3}, {4}, {5}}, {{1, 2, 3, 4}, {5}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getArraySubsets(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArraySubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubsets(t *testing.T) {
	type args struct {
		input []interface{}
	}
	tests := []struct {
		name string
		args args
		want [][][]interface{}
	}{
		{name: "two",
			args: args{[]interface{}{1, 2}},
			want: [][][]interface{}{{{1}, {2}}},
		},
		{name: "three",
			args: args{[]interface{}{1, 2, 3}},
			want: [][][]interface{}{{{1}, {2, 3}}, {{1}, {2}, {3}}, {{1, 2}, {3}}},
		},
		{name: "five",
			args: args{[]interface{}{1, 2, 3, 4, 5}},
			want: [][][]interface{}{{{1}, {2, 3, 4, 5}}, {{1}, {2}, {3, 4, 5}},
				{{1}, {2}, {3}, {4, 5}}, {{1}, {2}, {3}, {4}, {5}}, {{1}, {2}, {3, 4}, {5}}, {{1}, {2, 3}, {4, 5}},
				{{1}, {2, 3}, {4}, {5}}, {{1}, {2, 3, 4}, {5}}, {{1, 2}, {3, 4, 5}}, {{1, 2}, {3}, {4, 5}},
				{{1, 2}, {3}, {4}, {5}}, {{1, 2}, {3, 4}, {5}}, {{1, 2, 3}, {4, 5}}, {{1, 2, 3}, {4}, {5}}, {{1, 2, 3, 4}, {5}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subsets(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
