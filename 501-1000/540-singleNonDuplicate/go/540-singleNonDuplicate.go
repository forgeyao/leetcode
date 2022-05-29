// https://leetcode-cn.com/problems/single-element-in-a-sorted-array/
package main

import "fmt"

/**
 * 常规思路
 * 时间O(n/2), 即 O(n),不满足条件
 * 空间O(1)
 */
func singleNonDuplicate(nums []int) int {
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	return nums[len(nums)-1]
}

/**
 * 要时间 O(logn), 很容易想到二分法
 * 时间O(logn)
 * 空间O(1)
 */
func singleNonDuplicate2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (high + low) / 2
		/*if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				low = mid + 1
			} else {
				high = mid
			}
		} else {
			if nums[mid-1] == nums[mid] {
				low = mid + 1
			} else {
				high = mid
			}
		}*/
		// 异或特性
		if nums[mid] == nums[mid^1] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

/**
 * 参考官方解答. 二分法进一步优化
 * 时间O(logn)
 * 空间O(1)
 */
func singleNonDuplicate3(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := (high + low) / 2
		mid -= mid & 1
		if nums[mid] == nums[mid+1] {
			low = mid + 2
		} else {
			high = mid
		}
	}
	return nums[low]
}

func main() {
	nums := [][]int{{1, 1, 2, 3, 3, 4, 4, 8, 8}, {3, 3, 7, 7, 10, 11, 11}, {1, 1, 2}, {1, 1, 2, 3, 3}}
	for _, num := range nums {
		fmt.Println(num, "ret:", singleNonDuplicate(num), singleNonDuplicate2(num), singleNonDuplicate3(num))
	}
}
