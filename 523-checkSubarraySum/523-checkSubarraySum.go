// https://leetcode-cn.com/problems/continuous-subarray-sum/
package main

import (
	"fmt"
	"time"
)

func checkSubarraySum(nums []int, k int) bool {
	l := len(nums)
	prenum := 0
	mapnums := map[int]int{}
	mapnums[0] = -1
	for i := 0; i < l; i++ {
		prenum = prenum + nums[i]
		val, ok := mapnums[prenum%k]
		if ok {
			if i-val >= 2 {
				return true
			}
		} else {
			mapnums[prenum%k] = i
		}
	}
	return false
}
func checkSubarraySum2(nums []int, k int) bool {
	l := len(nums)
	prenums := make([]int, l)
	prenums[0] = nums[0]
	for i := 1; i < l; i++ {
		prenums[i] = prenums[i-1] + nums[i]
	}
	mapnums := make(map[int]int, 1)
	mapnums[0] = -1
	for i := 0; i < l; i++ {
		val, ok := mapnums[prenums[i]%k]
		if ok {
			if i-val >= 2 {
				return true
			}
		} else {
			mapnums[prenums[i]%k] = i
		}
	}
	return false
}
func checkSubarraySum3(nums []int, k int) bool {
	l := len(nums)
	for i := 0; i < l; i++ {
		sum := nums[i]
		for j := i + 1; j < l; j++ {
			sum = sum + nums[j]
			if sum%k == 0 {
				return true
			}
		}
	}
	return false
}

func main() {
	begin := time.Now().UnixNano() / int64(time.Millisecond)
	nums := [][]int{{23, 2, 4, 6, 7}, {23, 2, 6, 4, 7}, {23, 2, 4, 6, 6}, {1, 0}, {5, 0, 0, 0}, {0, 0}}
	k := []int{6, 6, 7, 2, 3, 1}
	for i := 0; i < len(nums); i++ {
		fmt.Println("ret:", checkSubarraySum(nums[i], k[i]))
	}
	end := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println("time:", end-begin, "ms")
	begin = end
	for i := 0; i < len(nums); i++ {
		fmt.Println("ret:", checkSubarraySum2(nums[i], k[i]))
	}
	end = time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Println("time:", end-begin, "ms")
}
