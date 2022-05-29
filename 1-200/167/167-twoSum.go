// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	var num []int
	for i, j := 0, len(numbers)-1; i < j; {
		if numbers[i]+numbers[j] == target {
			num = append(num, i+1, j+1)
			break
		}
		if numbers[i]+numbers[j] > target {
			j--
		} else {
			i++
		}
	}
	return num
}

func main() {
	n := [][]int{{2, 7, 11, 15}, {2, 3, 4}, {-1, 0}}
	target := []int{9, 6, -1}
	for i := 0; i < len(n) && i < len(target); i++ {
		fmt.Println("ret:", twoSum(n[i], target[i]))
	}
}
