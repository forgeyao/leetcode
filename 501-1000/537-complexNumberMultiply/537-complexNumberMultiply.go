// https://leetcode-cn.com/problems/complex-number-multiplication/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 思路：将实数、虚数从字符串里分离出来计算，最后再合并
 * 时间O(1)
 * 空间O(1)
 */
func complexNumberMultiply(num1 string, num2 string) string {
	a1, b1 := parse(num1)
	a2, b2 := parse(num2)
	a3 := a1*a2 - b1*b2
	b3 := a1*b2 + b1*a2
	return strconv.Itoa(a3) + "+" + strconv.Itoa(b3) + "i"
}

func parse(num string) (a, b int) {
	i := strings.IndexByte(num, '+')
	a, _ = strconv.Atoi(string(num[:i]))
	b, _ = strconv.Atoi(string(num[i+1 : len(num)-1]))
	return
}

func main() {
	num1s := []string{"1+1i", "1+-1i"}
	num2s := []string{"1+1i", "1+-1i"}
	for i := 0; i < len(num1s) && i < len(num2s); i++ {
		fmt.Println(complexNumberMultiply(num1s[i], num2s[i]))
	}
}
