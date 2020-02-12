//测试包注释
package main

import (
	"bytes"
	"fmt"
)

//测试注释
func main() {

	s := "你好!"
	fmt.Println(len(s))
	r := bytes.Runes([]byte(s))
	fmt.Println(len(r))
}
