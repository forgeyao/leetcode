package main

import (
	"fmt"
)

/**
 * 没想到特别的方法，就先用常规两层遍历实现，然后在此基础上做优化
 *
 * 1. 因为是绝对值，正/反是一样的，所以可以砍掉近一般的遍历
 * 2. 因为对数组下标有限制，可以进一步在循环里做限制，减少计算量
 * 时间复杂度 O(n*min(k,n))
 * 空间复杂的 O(1)
 */
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := i + 1; j <= k+i && j < l; j++ {
			if abs(nums[i]-nums[j]) <= t && abs(i-j) <= k {
				return true
			}
		}
	}
	return false
}

/**
 * 参考官方解答
 * 用桶缓存数据
 * 将 t+1 内的元素，即 [0,t] hash 到同一个桶
 * 桶内元素最多 k 个(超过就删掉最老的)
 *
 * 时间复杂度 O(n)
 * 空间复杂的 O(min(n,k))
 */
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}
func containsNearbyAlmostDuplicate2(nums []int, k int, t int) bool {
	m := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := m[id]; has {
			return true
		}
		if y, has := m[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := m[id+1]; has && abs(x-y) <= t {
			return true
		}
		m[id] = x
		if i >= k {
			delete(m, getID(nums[i-k], t+1))
		}
	}
	return false
}

func main() {
	nums := [][]int{{1, 2, 3, 1}, {1, 0, 1, 1}, {1, 5, 9, 1, 5, 9}}
	k := []int{3, 1, 2}
	t := []int{0, 2, 3}
	for i := 0; i < len(nums) && i < len(k) && i < len(t); i++ {
		fmt.Println("ret:", containsNearbyAlmostDuplicate(nums[i], k[i], t[i]), containsNearbyAlmostDuplicate2(nums[i], k[i], t[i]))
	}
}
