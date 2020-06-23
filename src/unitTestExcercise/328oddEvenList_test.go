package unitTestExcercise

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"1", args{create328()}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func create328() *ListNode {
	a5 := &ListNode{5, nil}
	a4 := &ListNode{4, a5}
	a3 := &ListNode{3, a4}
	a2 := &ListNode{2, a3}
	a1 := &ListNode{1, a2}
	return a1
}
