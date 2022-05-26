// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
package main

import (
	"fmt"
	"leetcode/util"
)

/**
 * 跟求解最大连续子串和类似
 *
 * 时间 O(n*m), n 为s长度, m 为最长子串长度
 * 空间 O(1)
 */
func lengthOfLongestSubstring(s string) int {
	max, start := 0, 0 // max 表示最大长度，start 当前子串起始位置
	for i := 0; i < len(s); i++ {
		for j := start; j < i; j++ {
			if s[j] == s[i] {
				start = j + 1
			}
		}
		if i-start+1 > max {
			max = i - start + 1
		}
	}
	return max
}

/**
 * 用 map 做缓存判断重复
 *
 * 时间 O(n), n 为s长度, m 为最长子串长度
 * 空间 O(∣Σ∣)，其中 Σ表示字符集个数
 */
func lengthOfLongestSubstring2(s string) int {
	m := map[byte]int{}
	max, start := 0, 0 // max 表示最大长度，start 当前子串起始位置
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok {
			if v >= start { // 没做删除, 所以这里需要判断
				start = v + 1
			}
		}
		m[s[i]] = i
		if i-start+1 > max {
			max = i - start + 1
		}
	}
	return max
}

/**
 * 官方答案。有点复杂，不容易理解
 *
 * 时间 O(n), n 为s长度
 * 空间 O(∣Σ∣)，其中 Σ表示字符集
 */

func lengthOfLongestSubstring3(s string) int {
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = util.Max(ans, rk-i+1)
	}
	return ans
}

func main() {
	ss := []string{"abcabcbb", "bbbbb", "pwwkew", "", " "}
	for _, s := range ss {
		fmt.Println("ret:", lengthOfLongestSubstring(s), lengthOfLongestSubstring2(s), lengthOfLongestSubstring3(s))
	}
}
