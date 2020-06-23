package unitTestExcercise

import "math"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	re := math.MaxInt64
	helper127(beginWord, endWord, wordList, 1, &re)
	if re == math.MaxInt64 {
		return 0
	}
	return re
}
func helper127(beginWord string, endWord string, wordList []string, cur int, re *int) {
	if beginWord == endWord {
		*re = min(*re, cur)
		return
	}
	for i := 0; i < len(beginWord); i++ {
		for index, elem := range wordList {
			j := 0
			for ; j < len(elem); j++ {
				if j != i && elem[j] != beginWord[j] {
					break
				}
			}
			if j == len(elem) {
				tmp := append([]string{}, wordList[:index]...)
				tmp = append(tmp, wordList[index+1:]...)
				helper127(elem, endWord, tmp, cur+1, re)
			}
		}
	}
}
