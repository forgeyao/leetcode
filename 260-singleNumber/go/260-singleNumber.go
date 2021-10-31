// https://leetcode-cn.com/problems/single-number-iii/
package main

import "fmt"

/**
 * 能快速想到的是用 map 做缓存
 * 时间复杂度 O(n)
 * 空间复杂度 O(n)
 */
func singleNumber(nums []int) []int {
	m := map[int]int{}
	for _, num := range nums {
		m[num]++
	}

	ret := []int{}
	for key, val := range m {
		if val == 1 {
			ret = append(ret, key)
		}
	}
	return ret
}

/**
 * 方法二 位运算
 * 有想到异或，但具体操作还是参考官方解答
 * 时间复杂度 O(n)
 * 空间复杂度 O(1)
 */

func singleNumber2(nums []int) []int {
	num1 := 0
	for _, num := range nums {
		num1 ^= num
	}

	lowOne := num1 & -num1
	num2, num3 := 0, 0
	for _, num := range nums {
		if num&lowOne > 0 {
			num2 ^= num
		} else {
			num3 ^= num
		}
	}
	return []int{num2, num3}
}

func main() {
	nums := [][]int{{1, 2, 1, 3, 2, 5}, {-1, 0}}
	for _, num := range nums {
		fmt.Println("ret:", singleNumber(num), singleNumber2(num))
	}
	x := 4
	fmt.Printf("%d %b %d %b %d %b\n", x, x, -x, -x, x&-x, x&-x)
	x = 7
	fmt.Printf("%d %b %d %b %d %b\n", x, x, -x, -x, x&-x, x&-x)
}
