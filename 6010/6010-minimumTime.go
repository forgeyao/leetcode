// https://leetcode-cn.com/problems/minimum-time-to-complete-trips/
package main

import (
	"fmt"
	"sort"
)

/**
 * 最开始思路是用数学办法，把所有数相乘求和，但这个导致越界
 * 然后该用二分查找
 * 时间 O(N * logN), N 为 time 长度
 * 空间 O(1)
 */
func minimumTime(time []int, totalTrips int) int64 {
	min := time[0]
	for _, v := range time {
		if v < min {
			min = v
		}
	}

	left, right := 0, totalTrips*min // 跑得最快的作为上限
	for left < right {
		mid, sum := left+(right-left)/2, 0
		for _, v := range time {
			sum += mid / v
		}
		if sum < totalTrips {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return int64(left)
}

/**
 * 抄的的一个很简洁的答案
 */
func minimumTime2(time []int, totalTrips int) int64 {
	return int64(sort.Search(totalTrips*1e7, func(x int) bool {
		tot := 0
		for _, t := range time {
			tot += x / t
			if tot >= totalTrips {
				return true
			}
		}
		return false
	}))
}

func main() {
	times := [][]int{
		{1, 2, 3},
		{2},
		{5, 10, 10},
		{39, 82, 69, 37, 78, 14, 93, 36, 66, 61, 13, 58, 57, 12, 70, 14, 67, 75, 91, 1, 34, 68, 73, 50, 13, 40, 81, 21, 79, 12, 35, 18, 71, 43, 5, 50, 37, 16, 15, 6, 61, 7, 87, 43, 27, 62, 95, 45, 82, 100, 15, 74, 33, 95, 38, 88, 91, 47, 22, 82, 51, 19, 10, 24, 87, 38, 5, 91, 10, 36, 56, 86, 48, 92, 10, 26, 63, 2, 50, 88, 9, 83, 20, 42, 59, 55, 8, 15, 48, 25},
	}
	totalTrips := []int{5, 1, 9, 4187}
	for i := 0; i < len(times) && i < len(totalTrips); i++ {
		fmt.Println(minimumTime(times[i], totalTrips[i]), minimumTime2(times[i], totalTrips[i]))
	}
}
