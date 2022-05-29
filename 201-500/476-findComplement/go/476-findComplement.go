// https://leetcode-cn.com/problems/number-complement/
package main

import (
	"fmt"
)

func findComplement(num int) int {
	ret := 0
	for n, m, i := num/2, num%2, 1; n != 0; n, m, i = n/2, n%2, i*2 {
		if m == 0 {
			ret += i
		}
	}
	return ret
}

func findComplement2(num int) int {
	highBit := 0
	// 找最高为1的位
	for i := 1; i <= 30; i++ {
		if num < 1<<i {
			break
		}
		highBit = i
	}
	mask := 1<<(highBit+1) - 1
	return num ^ mask
}

func main() {
	num := []int{5, 1, 4}
	for _, n := range num {
		fmt.Println(n, " ret:", findComplement(n), findComplement2(n))
	}
}
