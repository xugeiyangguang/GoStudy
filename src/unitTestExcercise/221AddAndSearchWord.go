package unitTestExcercise

type WordDictionary struct {
	dict map[int][]string
}

/** Initialize your data structure here. */
func Constructor173() WordDictionary {
	return WordDictionary{make(map[int][]string, 0)}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	if this.dict[len(word)] != nil {
		this.dict[len(word)] = append(this.dict[len(word)], word)
	} else {
		this.dict[len(word)] = []string{word}
	}
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	if this.dict[len(word)] == nil {
		return false
	}
	for i := range this.dict[len(word)] {
		tag := true
		for j := 0; j < len(word); j++ {
			if word[j] != '.' && this.dict[len(word)][i][j] != word[j] {
				tag = false
			}
		}
		if tag {
			return true
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
