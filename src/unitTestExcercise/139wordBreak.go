package unitTestExcercise

import (
	"math"
)

func wordBreak(s string, wordDict []string) bool {
	mapp := make(map[string]int, 0)
	min := math.MaxInt64
	max := 0
	for i := range wordDict {
		mapp[wordDict[i]] = len(wordDict[i])
		if len(wordDict[i]) < min {
			min = len(wordDict[i])
		}
		if len(wordDict[i]) > max {
			max = len(wordDict[i])
		}
	}
	return helper1(s, 0, mapp, min, max)
}

func helper1(s string, start int, mapp map[string]int, min, max int) bool {
	if start == len(s) {
		return true
	}
	if len(s)-start < min {
		return false
	}
	for i := start; i < len(s); i++ {
		if i-start+1 < min {
			continue
		}
		if i-start+1 > max {
			return false
		}
		tmp := s[start : i+1]
		if mapp[tmp] != 0 {
			re := helper1(s, i+1, mapp, min, max)
			if re == true {
				return true
			}
		}
	}
	return false
}
