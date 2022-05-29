// https://leetcode.cn/problems/number-of-1-bits/
package main

import "fmt"

/**
 * 利用二进制特点
 *
 * 时间 O(k)，其中 k 是 int 型的二进制位数
 * 空间 O(1)
 */
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num%2 == 1 {
			count++
		}
		num /= 2
	}
	return count
}

/**
 * 参考答案。n&(n-1) 可以把地位1清0，加速计算
 *
 * 时间 O(logN)，其中 n 是 二进制1的个数
 * 空间 O(1)
 */
func hammingWeight2(num uint32) int {
	count := 0
	for ; num > 0; num &= num - 1 {
		count++
	}
	return count
}

func main() {
	nums := []uint32{00000000000000000000000000001011, 00000000000000000000000010000000}
	for _, num := range nums {
		fmt.Println("ret:", hammingWeight(num), hammingWeight2(num))
	}
}
