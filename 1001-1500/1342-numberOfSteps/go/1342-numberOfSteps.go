// https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/
package main

import "fmt"

func numberOfSteps(num int) int {
	count := 0
	for num != 0 {
		if num%2 == 0 {
			num >>= 1
		} else {
			num -= 1
		}
		count++
	}
	return count
}

func main() {
	nums := []int{14, 8}
	for _, num := range nums {
		fmt.Println("ret:", numberOfSteps(num))
	}
}
