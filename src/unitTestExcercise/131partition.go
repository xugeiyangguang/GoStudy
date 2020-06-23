package unitTestExcercise

func partition1(s string) [][]string {
	re := [][]string{}
	if len(s) == 0 {
		return re
	}
	helper7(s, []string{}, &re)
	return re
}

func helper7(s string, cur []string, re *[][]string) {
	if len(s) == 0 {
		*re = append(*re, append([]string{}, cur...))
		return
	}
	for i := 1; i <= len(s); i++ {
		tmp := s[0:i]
		if stringIsPalindrome(tmp) {
			helper7(s[i:], append(cur, tmp), re)
		}
	}
}

func stringIsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

//func main(){
//	s:="aab"
//	re:=partition(s)
//	fmt.Println(re)
//}
