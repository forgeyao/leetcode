// https://leetcode-cn.com/problems/valid-perfect-square/
package main

import "fmt"

/**
 * 二分法查找是否存在合适的值 n
 * n*n == num
 */
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	low, mid, high := 1, num/2, num-1
	for low <= mid && mid <= high {
		square := mid * mid
		if square == num {
			return true
		} else if square < num {
			low = mid + 1
		} else {
			high = mid - 1
		}
		mid = (low + high) / 2
	}
	return false
}

func main() {
	num := []int{16, 14}
	for _, n := range num {
		fmt.Println("ret:", isPerfectSquare(n))
	}
}
