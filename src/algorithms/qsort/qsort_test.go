package qsort

import "testing"

func TestQuickSort1(t *testing.T) {
	Values := []int{4, 3, 2, 1}
	QuickSort(Values)
	if Values[0] != 1 || Values[1] != 2 || Values[2] != 3 || Values[3] != 4 {
		t.Errorf("快速排序失败！")
	}
}

func TestQuickSort2(t *testing.T) {
	Values := []int{5, 5, 3, 2, 1}
	QuickSort(Values)
	if Values[0] != 1 || Values[1] != 2 || Values[2] != 3 || Values[3] != 5 || Values[4] != 5 {
		t.Errorf("快速排序失败！")
	}
}
