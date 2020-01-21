package main

import "fmt"

func main() {

	s := "abcde"
	for i := 0; i < len(s); i++ { //byte=uint8    代表了 ASCII 码的一个字符
		tmp := s[i]
		fmt.Println(tmp)
	}

	for _, v := range s { //rune=int32     代表一个 UTF-8 字符
		tmp := v
		fmt.Println(tmp)
	}
}
