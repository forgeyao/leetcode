// https://leetcode-cn.com/problems/pancake-sorting/
package main

import "fmt"

/**
 * 思路：由于只能翻转前面的子数组。根据这个特性可以用冒泡排序
 * 从结尾开始，每次把最大值排到最后，直到首位
 * 时间O(n*n)
 * 空间O(1)
 */
func pancakeSort(arr []int) []int {
	ret := []int{}
	for i := len(arr) - 1; i > 0; i-- {
		maxIdx := i
		for j := 0; j < i; j++ { // 找最大值
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		if maxIdx != i { // 最大值不在原位
			if maxIdx > 0 { // 最大值在首位，可以省略此次翻转
				overturn(arr, maxIdx) // 把最大值翻到最前
				ret = append(ret, maxIdx+1)
			}
			overturn(arr, i) // 再把最大值翻到最后
			ret = append(ret, i+1)
		}
	}
	return ret
}

func overturn(arr []int, maxIdx int) {
	for i, j := 0, maxIdx; i < j; i++ {
		arr[i], arr[j] = arr[j], arr[i]
		j--
	}
}

func main() {
	arrs := [][]int{
		{3, 2, 4, 1},
		{1, 2, 3},
	}
	for _, arr := range arrs {
		fmt.Println(pancakeSort(arr))
	}
}
