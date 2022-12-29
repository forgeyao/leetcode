// https://leetcode.cn/problems/valid-anagram/
package main

import "fmt"

/**
 * map 缓存统计情况，再遍历另一个做对比
 * 时间 O(N)
 * 空间 O(N)
 */
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms := map[byte]int{}
	for i := 0; i < len(s); i++ {
		ms[s[i]]++
	}
	for i := 0; i < len(t); i++ {
		ms[t[i]]--
		if ms[t[i]] < 0 {
			return false
		}
	}
	return true
}

func main() {
	ss := []string{"anagram", "rat", "a"}
	tt := []string{"nagaram", "car", "ab"}

	for i := 0; i < len(ss) && i < len(tt); i++ {
		fmt.Println(isAnagram(ss[i], tt[i]))
	}
}
