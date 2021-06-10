// https://leetcode-cn.com/problems/climbing-stairs/
package main

import "fmt"

// 递归 斐波拉契数,会超时
func climbStairs2(n int) int {
	if n == 0 {
		return 1
	}
	num := 0
	if n-1 >= 0 {
		num += climbStairs(n - 1)
	}
	if n-2 >= 0 {
		num += climbStairs(n - 2)
	}
	return num
}

// 动态规划
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	n1, n2 := 1, 2
	num := 0
	for i := 3; i <= n; i++ {
		num = n1 + n2
		n1 = n2
		n2 = num
	}
	return num
}

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 44}
	for i := 0; i < len(n); i++ {
		fmt.Println("ret:", climbStairs(n[i]))
	}
}
