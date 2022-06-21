// https://leetcode.cn/problems/kth-largest-element-in-an-array/
package main

import (
	"fmt"
	"sort"
)

/**
 * 常规办法. 先排序
 * 时间O(N*logN)
 * 空间O(1)
 */
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

/**
 * 想到堆排序。忘了怎么构造，结果就实现成 Top k 排序了
 * 时间O(N*k*k)
 * 空间O(k)
 */
func findKthLargest2(nums []int, k int) int {
	heap := []int{}
	for i := 0; i < k; i++ {
		heap = append(heap, -10001)
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < k; j++ {
			if nums[i] > heap[j] {
				for m := k - 1; m > j; m-- {
					heap[m] = heap[m-1]
				}
				heap[j] = nums[i]
				break
			}
		}
		//fmt.Println("heap:", heap)
	}
	return heap[len(heap)-1]
}

/**
 * 大顶堆,堆排序
 */
func findKthLargest3(nums []int, k int) int {
	l := len(nums)
	// 调整堆
	for i := l/2 - 1; i >= 0; i-- {
		adustHeap(nums, i, l)
	}

	for i := 0; i < k; i++ {
		nums[0], nums[l-i-1] = nums[l-i-1], nums[0] // 堆顶和最后接一个交换
		adustHeap(nums, 0, l-i-1)
	}
	return nums[l-k]
}

func adustHeap(nums []int, i int, l int) {
	left := 2*i + 1
	if left > l-1 {
		return
	}
	if nums[i] < nums[left] {
		nums[i], nums[left] = nums[left], nums[i]
		adustHeap(nums, left, l)
	}
	if left+1 < l && nums[i] < nums[left+1] {
		nums[i], nums[left+1] = nums[left+1], nums[i]
		adustHeap(nums, left+1, l)
	}
}

/**
 * 大顶堆,堆排序。参考答案优化 findKthLargest3
 * 时间O(N*logN), 建堆O(n),删除O(k*logN)
 * 空间O(logN)
 */
func findKthLargest4(nums []int, k int) int {
	l := len(nums)
	// 建堆
	for i := l/2 - 1; i >= 0; i-- {
		adustHeap2(nums, i, l)
	}

	// 移除 k-1 个顶，返回堆顶就行
	for i := 0; i < k-1; i++ {
		nums[0], nums[l-i-1] = nums[l-i-1], nums[0] // 堆顶和最后接一个交换
		adustHeap2(nums, 0, l-i-1)
	}
	return nums[0]
}

func adustHeap2(nums []int, i int, l int) {
	left, right, max := 2*i+1, 2*i+2, i
	if left < l && nums[max] < nums[left] {
		max = left
	}
	if right < l && nums[max] < nums[right] {
		max = right
	}
	if max != i {
		nums[i], nums[max] = nums[max], nums[i]
		adustHeap2(nums, max, l)
	}
}

func main() {
	nums := [][]int{
		{3, 2, 1, 5, 6, 4},
		{3, 2, 3, 1, 2, 4, 5, 5, 6},
	}
	k := []int{2, 4}
	result := []int{5, 4}
	for i := 0; i < len(nums) && i < len(k) && i < len(result); i++ {
		//fmt.Println(nums[i], k[i], " ret:", findKthLargest(nums[i], k[i]), "result:", result[i])
		//fmt.Println(nums[i], k[i], " ret:", findKthLargest2(nums[i], k[i]), "result:", result[i])
		//fmt.Println(nums[i], k[i], " ret:", findKthLargest3(nums[i], k[i]), "result:", result[i])
		fmt.Println(nums[i], k[i], " ret:", findKthLargest4(nums[i], k[i]), "result:", result[i])
	}
}
