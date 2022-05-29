// https://leetcode-cn.com/problems/buddy-strings/
package main

import "fmt"

/**
 * 根据字符串不同长度做判断
 * 1. 小于2，false
 * 2. 大于等于2，最多只允许2个位置字符不同
 *
 * 时间复杂度 O(n)
 * 空间复杂度 O(n) (n <= 26)
 */
func buddyStrings(s string, goal string) bool {
	l1, l2 := len(s), len(goal)
	if l1 < 2 || l1 != l2 {
		return false
	}

	diff := []int{}      // 存不同的字符
	m := map[byte]bool{} // 存不同字符，解决两个字符串完全相同下，字符串内是否存在相同字符
	for i := 0; i < l1; i++ {
		m[s[i]] = true
		if s[i] != goal[i] {
			diff = append(diff, i)
		}
		if len(diff) > 2 {
			return false
		}
	}
	if len(diff) == 0 {
		if len(m) == l1 {
			return false
		}
		return true
	} else if len(diff) == 1 {
		return false
	}
	return s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
}

func main() {
	s := []string{"ab", "ab", "aa", "aaaaaaabc", "abc", "abcd"}
	goal := []string{"ba", "ab", "aa", "aaaaaaacb", "acd", "abcd"}
	for i := 0; i < len(s) && i < len(goal); i++ {
		fmt.Println(s[i], goal[i], buddyStrings(s[i], goal[i]))
	}
}
