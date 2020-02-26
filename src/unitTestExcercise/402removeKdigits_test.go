package unitTestExcercise

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"1432219", 3}, "1219"},
		{"2", args{"10200", 1}, "200"},
		{"3", args{"10", 1}, "0"},
		{"4", args{"112", 1}, "11"},
		{"5", args{"1111111", 3}, "1111"},
		{"6", args{"61023", 2}, "23"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
