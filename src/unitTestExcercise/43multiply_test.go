package unitTestExcercise

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"123", "456"}, "56088"},
		{"2", args{"9", "9"}, "81"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_single(t *testing.T) {
	type args struct {
		num string
		a   byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"123", '6'}, "738"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := single(tt.args.num, tt.args.a); got != tt.want {
				t.Errorf("single() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_strAdd(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"1", args{"111", "222"}, "333"},
		{"2", args{"738", ""}, "738"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("strAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
