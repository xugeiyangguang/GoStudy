package unitTestExcercise

func generateParenthesis1(n int) []string {
	if n == 0 {
		return []string{}
	}
	re := make([]string, 0)
	helper22(n, n, "", &re)
	return re
}

func helper22(left, right int, cur string, re *[]string) {
	if left == 0 && right == 0 {
		*re = append(*re, cur)
		return
	}

	if left > 0 {
		helper22(left-1, right, cur+"(", re)
	}
	if left < right && right > 0 {
		helper22(left, right-1, cur+")", re)
	}

}
