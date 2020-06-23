package unitTestExcercise

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{creat337()}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.root); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func creat337() *TreeNode {
	a5 := &TreeNode{1, nil, nil}
	a4 := &TreeNode{3, nil, nil}
	a3 := &TreeNode{3, nil, a5}
	a2 := &TreeNode{2, nil, a4}
	a1 := &TreeNode{3, a2, a3}
	return a1
}
