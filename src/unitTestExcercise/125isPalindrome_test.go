package unitTestExcercise

import "testing"

func Test_isPalindrome1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{"0P"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome1(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome1() = %v, want %v", got, tt.want)
			}
		})
	}
}
