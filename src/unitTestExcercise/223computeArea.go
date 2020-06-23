package unitTestExcercise

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	s1 := (C - A) * (D - B)
	s2 := (G - E) * (H - F)
	a := max(A, E)
	b := max(B, F)
	c := min(C, G)
	d := min(D, H)
	S_common := 0
	if c-a > 0 && d-b > 0 {
		S_common = (c - a) * (d - b)
	}
	re := s1 + s2 - S_common
	return re
}
