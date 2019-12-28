package Simplemath

import "fmt"

type Stack struct {
	arr []interface{}
}

func (s *Stack) CreateStack() *Stack {
	s.arr = []interface{}{}
	return s
}

func (s *Stack) push(a interface{}) {
	s.arr = append(s.arr, a)
}

func (s *Stack) pop() (a interface{}, error error) {
	if len(s.arr) == 0 {
		//fmt.Errorf可以打印错误信息
		error = fmt.Errorf("栈中没有元素")
	}
	a = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return
	//return a,error
}
