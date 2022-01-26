package main

import (
	"reflect"
	"testing"
)

func Test_letterCombinationsOfPhoneNumber(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "56",
			args: args{digits: "56"},
			want: []string{"jm", "jn", "jo", "km", "kn", "ko", "lm", "ln", "lo"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinationsOfPhoneNumber(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinationsOfPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
