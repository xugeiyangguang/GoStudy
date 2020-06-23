package src

import (
	"fmt"
	"strconv"
)

func reverse(x int) int {

	flag := 'a'

	str := strconv.Itoa(x)
	if str[0] < '0' || str[0] > '9' {
		flag = int32(str[0])
		str = str[1:]
	}
	str1 := ""
	for i := len(str) - 1; i >= 0; i-- {
		str1 += string(str[i])
	}

	if flag != 'a' {
		str1 = string(flag) + str1
	}
	result, error := strconv.Atoi(str1)
	if error != nil {
		return 0
	}

	return result

}

func main() {
	re := reverse(20)
	fmt.Print(re)
}
