package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 {
		t.Errorf("冒泡排序失败！")
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Errorf("冒泡排序失败！")
	}
}
