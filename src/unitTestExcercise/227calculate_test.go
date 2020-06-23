package unitTestExcercise

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		//	{"1",args{"3+2*2"},7},
		//	{"2",args{"3+2*10"},23},
		//	{"3",args{" 3/2 "},1},
		{"3", args{"2-3+4"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
