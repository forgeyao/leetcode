// https://leetcode.cn/problems/sort-an-array/
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/**
 * 快排
 * 时间 O(NlogN)
 * 空间 O(1)
 */
func sortArray(nums []int) []int {
	sortPart(0, len(nums)-1, nums)
	return nums
}

func sortPart(l, r int, nums []int) {
	if l >= r {
		return
	}
	pivot := rand.Intn(r-l+1) + l
	nums[pivot], nums[r] = nums[r], nums[pivot]
	pivot = r

	small := l - 1
	for big := l; big < r; big++ {
		if nums[big] <= nums[pivot] {
			small += 1
			nums[small], nums[big] = nums[big], nums[small]
		}
	}
	nums[small+1], nums[pivot] = nums[pivot], nums[small+1]
	pivot = small + 1

	sortPart(l, pivot-1, nums)
	sortPart(pivot+1, r, nums)
}

/**
 * 归并排序
 * 时间 O(NlogN)
 * 空间 O(N)
 */
func sortArray2(nums []int) []int {
	sortInner(0, len(nums)-1, nums)
	return nums
}

func sortInner(l, r int, nums []int) {
	le := r - l + 1
	if le <= 1 {
		return
	}
	if le == 2 {
		if nums[l] > nums[r] {
			nums[l], nums[r] = nums[r], nums[l]
		}
		return
	}
	mid := (r-l)/2 + l
	sortInner(l, mid, nums)
	sortInner(mid+1, r, nums)
	merge(l, mid, r, nums)
}

func merge(l, mid, r int, nums []int) {
	// 有的是新申请临时数组来做合并
	for i := mid + 1; i <= r; i++ {
		pos := sort.SearchInts(nums[l:i], nums[i]) + l
		tmp := nums[i]
		for j := i; j > pos; j-- {
			nums[j] = nums[j-1]
		}
		nums[pos] = tmp
	}
}

func main() {
	numss := [][]int{
		{1},
		{1, 2},
		{5, 2},
		{5, 2, 3, 1},
		{5, 1, 1, 2, 0, 0},
	}
	answers := [][]int{
		{1},
		{1, 2},
		{2, 5},
		{1, 2, 3, 5},
		{0, 0, 1, 1, 2, 5},
	}

	for i := 0; i < len(numss) && i < len(answers); i++ {
		tmp := numss[i]
		fmt.Println(tmp, " ans:", answers[i], "ret:", sortArray(tmp))
		tmp = numss[i]
		fmt.Println(tmp, " ans:", answers[i], "ret:", sortArray2(tmp))
	}
}
