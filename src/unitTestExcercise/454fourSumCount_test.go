package unitTestExcercise

import "testing"

func Test_fourSumCount(t *testing.T) {
	type args struct {
		A []int
		B []int
		C []int
		D []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{0}, []int{0}, []int{0}, []int{0}}, 1},
		{"2", args{[]int{0, 1, -1},
			[]int{-1, 1, 0},
			[]int{0, 0, 1},
			[]int{-1, 1, 1}}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSumCount(tt.args.A, tt.args.B, tt.args.C, tt.args.D); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
