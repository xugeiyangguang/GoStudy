package unitTestExcercise

import (
	"reflect"
	"testing"
)

func Test_findLast(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{5, 7, 7, 8, 8, 10}, 8}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLast(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findLast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findStart(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{5, 7, 7, 8, 8, 10}, 8}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findStart(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findStart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{"1", args{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
