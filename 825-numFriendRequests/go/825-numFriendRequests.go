// https://leetcode-cn.com/problems/friends-of-appropriate-ages/
package main

import (
	"fmt"
	"sort"
)

/**
 * 排序 + 双指针
 */
func numFriendRequests(ages []int) (count int) {
	sort.Ints(ages)
	left, right := 0, 0 // 表示 x
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ages[left]*2 <= age+14 {
			left++
		}
		for right+1 < len(ages) && ages[right+1] <= age {
			right++
		}
		count += right - left
	}
	return
}

/**
 * 计数排序 + 前缀和
 */
func numFriendRequests2(ages []int) (ans int) {
	const mx = 121
	var cnt, pre [mx]int
	for _, age := range ages {
		cnt[age]++
	}
	for i := 1; i < mx; i++ {
		pre[i] = pre[i-1] + cnt[i]
	}
	for i := 15; i < mx; i++ {
		if cnt[i] > 0 {
			bound := i/2 + 8
			ans += cnt[i] * (pre[i] - pre[bound-1] - 1)
		}
	}
	return
}

func main() {
	ages := [][]int{{16, 16}, {16, 17, 18}, {20, 30, 100, 110, 120}}
	for _, age := range ages {
		fmt.Println("ret:", numFriendRequests(age), numFriendRequests2(age))
	}
}
