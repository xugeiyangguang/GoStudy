package Simplemath

import "testing"

//可以对方法分开测试，
//也可以对方法进行整体测试，这个时候需要考虑整体的输出结果，包括前面的处理，不是只考虑方法中
var stack Stack

func TestStack_CreateStack(t *testing.T) {
	stack.CreateStack()
	if stack.arr == nil {
		t.Errorf("初始化失败")
	}
}

func TestStack_push(t *testing.T) {
	stack.push(1)
	stack.push(2)
	stack.push(3)

	if len(stack.arr) != 3 {
		t.Errorf("入栈失败！")
	}
}

func TestStack_Pop(t *testing.T) {
	stack.push(1)
	stack.push(1)
	stack.push(1)

	stack.pop()
	stack.pop()

	if len(stack.arr) != 3 {
		t.Errorf("出栈失败！")
	}
}
