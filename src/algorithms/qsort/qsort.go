package qsort

func QuickSort(in []int) {
	quickSort(in, 0, len(in)-1)
}

func quickSort(Values []int, left, right int) {
	key := Values[left]
	i, j := left, right
	for i < j {
		for i < j && Values[j] >= key {
			j--
		}
		Values[i] = Values[j]
		for i < j && Values[i] <= key {
			i++
		}
		Values[j] = Values[i]
	}
	Values[i] = key
	if i-left > 1 {
		quickSort(Values, left, i-1)
	}
	if right-i > 1 {
		quickSort(Values, i+1, right)
	}
}
