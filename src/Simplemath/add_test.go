package Simplemath

import "testing"

func TestAdd(t *testing.T) {
	re := Add(1, 1)
	if re != 3 {
		t.Errorf("Add 方法出错了！")
	}

}
