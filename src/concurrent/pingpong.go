package main

import (
	"fmt"
	"time"
)

func playera(table chan int) {
	name := "a"
	for {
		ball := <-table
		ball++
		fmt.Println("接球人：", name)
		time.Sleep(1 * time.Second)
		table <- ball
	}
}
func playerc(table chan int) {
	name := "c"
	for {
		ball := <-table
		ball++
		fmt.Println("接球人：", name)
		time.Sleep(1 * time.Second)
		table <- ball
	}
}
func playerb(table chan int) {

	name := "b"
	for {
		ball := <-table
		ball++
		fmt.Println("接球人：", name)
		time.Sleep(1 * time.Second)
		table <- ball
	}
}
func main() {
	var ball int
	table := make(chan int)
	go playera(table)
	go playerb(table)
	go playerc(table)
	table <- ball
	time.Sleep(10 * time.Second)
	<-table
}
