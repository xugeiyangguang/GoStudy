package main

import "fmt"

//const a = 1 << iota  //1  iota==0
//const b = 1 << iota  //1  iota==0
//const c  = 3         //3  iota==0
//const d  = 1<< iota  //1  iota==0

const (
	a = 1 << iota //1   iota==0
	b = 1 << iota //2   iota==1
	c = 3         //3   iota==2(unused)
	d = 1 << iota //8   iota==3
	e             //16  iota==4(复制使用上一次的表达式)
	f = 2 << iota //64  iota==5(表达式更换)
	g             //128 iota==6(复制使用新更换的表达式)
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
