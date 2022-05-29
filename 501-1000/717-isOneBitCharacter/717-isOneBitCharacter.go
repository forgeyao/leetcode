// https://leetcode-cn.com/problems/1-bit-and-2-bit-characters/solution/
package main

import "fmt"

/**
 * 思路：遍历一遍，如果是"1"就向前移两位, "0" 只需判断如果是最后一位就返回 true
 * 时间O(n)
 * 空间O(1)
 */
func isOneBitCharacter(bits []int) bool {
	for i := 0; i < len(bits); i++ {
		if bits[i] == 1 {
			i++
		} else {
			if i == len(bits)-1 {
				return true
			}
		}
	}
	return false
}

/**
 * 官方答案，更简洁
 */
func isOneBitCharacter2(bits []int) bool {
	i, n := 0, len(bits)
	for i < n-1 {
		i += bits[i] + 1
	}
	return i == n-1
}

func main() {
	bitss := [][]int{
		{1, 0, 0},
		{1, 1, 1, 0},
	}
	for _, bits := range bitss {
		fmt.Println(bits, isOneBitCharacter(bits))
	}
}
