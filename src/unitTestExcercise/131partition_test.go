package unitTestExcercise

import "testing"

func Test_stringIsPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{"aba"}, true},
		{"2", args{"aab"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringIsPalindrome(tt.args.s); got != tt.want {
				t.Errorf("stringIsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
