package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("C:\\Users\\27124\\Desktop\\neo4j.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content))

	err1 := ioutil.WriteFile("C:\\Users\\27124\\Desktop\\neo4jjjj.txt", content, 0666)
	if err1 != nil {
		fmt.Println(err1)
	}
}
