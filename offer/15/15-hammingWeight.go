// https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
package main

import "fmt"

func hammingWeight(num uint32) int {
	var base uint32 = 2
	count := 0
	if num%2 == 1 {
		count = 1
	}
	for i := 1; i < 32; i++ {
		if base&num > 0 {
			count++
		}
		base *= 2
	}
	return count
}

// 更简洁版
func hammingWeight2(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if (1<<i)&num > 0 {
			count++
		}
	}
	return count
}

// 更牛的版本
func hammingWeight3(num uint32) int {
	count := 0
	for ; num > 0; num &= (num - 1) {
		count++
	}
	return count
}

func main() {
	n := []uint32{00000000000000000000000000001011, 9, 2345}
	for i := 0; i < len(n); i++ {
		fmt.Println("num,ret:", n[i], hammingWeight(n[i]), hammingWeight2(n[i]), hammingWeight3(n[i]))
	}
}
