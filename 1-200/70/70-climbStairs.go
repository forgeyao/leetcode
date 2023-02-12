// https://leetcode.cn/problems/climbing-stairs/
package main

import "fmt"

/**
 * 递归 斐波拉契数,会超时
 * 时间 O(2n次方)
 * 空间 O(N)
 */
func climbStairs2(n int) int {
	if n <= 2 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

/**
 * 动态规划
 * 时间 O(N)
 * 空间 O(1)
 */
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	n1, n2 := 1, 2
	num := 0
	for i := 3; i <= n; i++ {
		num = n1 + n2
		n1, n2 = n2, num
	}
	return num
}

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 44}
	answers := []int{1, 2, 3, 5, 8, 13, 1134903170}
	for i := 0; i < len(n); i++ {
		fmt.Println(n[i], "ans:", answers[i], "ret:", climbStairs(n[i]))
	}
}
