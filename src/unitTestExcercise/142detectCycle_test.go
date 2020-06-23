package unitTestExcercise

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"1", args{init142()}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
func init142() *ListNode {
	a3 := &ListNode{3, nil}
	a2 := &ListNode{2, nil}
	a0 := &ListNode{0, nil}
	a4 := &ListNode{-4, nil}
	a3.Next = a2
	a2.Next = a0
	a0.Next = a4
	a4.Next = a2
	return a3
}
