// https://leetcode-cn.com/problems/sum-of-unique-elements/
package main

import "fmt"

/**
 * 常规办法. 用 map 做重复判断
 * 时间 O(n), n 是 nums 的长度
 * 空间 O(n), n 是哈希表需要的空间
 */
func sumOfUnique(nums []int) int {
	var m = map[int]int{}
	for _, num := range nums {
		m[num]++
	}
	sum := 0
	for k, v := range m {
		if v == 1 {
			sum += k
		}
	}
	return sum
}

/**
 * 在常规方法上做的细节优化. 增加状态记录，只需一次遍历
 * 时间 O(n), n 是 nums 的长度
 * 空间 O(n), n 是哈希表需要的空间
 */
func sumOfUnique2(nums []int) int {
	m := map[int]int{}
	sum := 0
	for _, num := range nums {
		if m[num] == 0 {
			sum += num
		} else if m[num] == 1 {
			sum -= num
		}
		m[num]++
	}
	return sum
}

func main() {
	numss := [][]int{
		{1, 2, 3, 2},
		{1, 1, 1, 1, 1},
		{1, 2, 3, 4, 5},
	}
	for _, nums := range numss {
		fmt.Println(nums, " ret:", sumOfUnique(nums), sumOfUnique2(nums))
	}
}
