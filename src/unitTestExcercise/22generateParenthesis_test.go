package unitTestExcercise

import (
	"reflect"
	"testing"
)

func Test_generateParenthesis1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"1", args{3}, []string{"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateParenthesis1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_helper22(t *testing.T) {
	type args struct {
		left  int
		right int
		cur   string
		re    []string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
