package unitTestExcercise

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{"226"}, 3},
		{"2", args{"12"}, 2},
		{"3", args{"0"}, 0},
		{"4", args{"101"}, 1},
		{"5", args{"10"}, 1},
		{"6", args{"110"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
