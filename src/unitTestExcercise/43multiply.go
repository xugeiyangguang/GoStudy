package unitTestExcercise

import "strconv"

func multiply(num1 string, num2 string) string {
	re := ""
	for i := len(num2) - 1; i >= 0; i-- {
		tmp := single(num1, num2[i])
		for k := 0; k < len(num2)-1-i; k++ {
			tmp += "0"
		}
		re = strAdd(re, tmp)
	}
	k, _ := strconv.Atoi(re)
	if k == 0 {
		return "0"
	}
	return re
}
func single(num string, a byte) string {
	add := 0
	re := ""
	for i := len(num) - 1; i >= 0; i-- {
		tmp := int(a-'0')*int(num[i]-'0') + add
		add = tmp / 10
		tmp = tmp % 10
		re = strconv.Itoa(tmp) + re
	}
	if add != 0 {
		re = strconv.Itoa(add) + re
	}
	return re
}
func strAdd(a, b string) string {
	if len(a) < len(b) {
		gap := len(b) - len(a)
		for k := 0; k < gap; k++ {
			a = "0" + a
		}
	}
	if len(b) < len(a) {
		gap := len(a) - len(b)
		for k := 0; k < gap; k++ {
			b = "0" + b
		}
	}
	add := 0
	re := ""
	i := len(a) - 1
	for ; i >= 0; i-- {
		tmp := int(a[i]-'0') + int(b[i]-'0') + add
		add = tmp / 10
		tmp = tmp % 10
		re = strconv.Itoa(tmp) + re
	}
	if add != 0 {
		re = strconv.Itoa(add) + re
	}
	return re
}
