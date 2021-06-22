// https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/
package main

import "fmt"

func permutation(s string) []string {
	l := len(s)
	if l == 1 {
		return []string{s}
	}

	var ret []string
	m := make(map[byte]int) // 字符去重
	for i := 0; i < l; i++ {
		if _, ok := m[s[i]]; ok {
			continue
		}
		ss := permutation(s[0:i] + s[i+1:]) // 递归
		for _, v := range ss {
			ret = append(ret, string(s[i])+v)
		}
		m[s[i]] = 1
	}
	return ret
}

func main() {
	s := []string{"a", "ab", "abc", "abcd", "aab"}
	for _, v := range s {
		fmt.Println("ret:", permutation(v))
	}
}
