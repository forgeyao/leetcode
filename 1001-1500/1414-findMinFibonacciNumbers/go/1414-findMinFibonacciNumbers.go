// https://leetcode-cn.com/problems/find-the-minimum-number-of-fibonacci-numbers-whose-sum-is-k/
package main

import "fmt"

/**
 * 缓存斐波拉切数列,其中最大值<=k
 * 倒序遍历数列,由大到小构造k
 * 时间：O(logk)
 * 空间: O(logk)
 */
func findMinFibonacciNumbers(k int) int {
	var fib = []int{1, 1}
	for fib[len(fib)-1] < k {
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
	}
	count := 0
	for i := len(fib) - 1; i >= 0 && k != 0; i-- {
		count += k / fib[i]
		k %= fib[i]
	}
	return count
}

func main() {
	ks := []int{7, 10, 19}
	for _, k := range ks {
		fmt.Println(k, findMinFibonacciNumbers(k))
	}
}
