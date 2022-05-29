// https://leetcode-cn.com/problems/implement-strstr/
package main

import "fmt"

// 暴力法
func strStr2(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		j, k := 0, i
		for ; j < len(needle) && k < len(haystack) && haystack[k] == needle[j]; j, k = j+1, k+1 {
		}
		if j == len(needle) && j == k-i {
			return i
		}
	}
	return -1
}

func strStr(haystack string, needle string) int {
	m := len(needle)
	if m == 0 {
		return 0
	}

	next := make([]int, m)
	next[0] = 0
	for i, j := 1, 0; i < m; i++ {
		if needle[i] == needle[j] {
			j++
		} else {
			for j > 0 && needle[i] != needle[j] {
				j = next[j-1]
			}
			if needle[i] == needle[j] {
				j++
			}
		}
		next[i] = j
	}

	n := len(haystack)
	for i, j := 0, 0; i < n; i++ {
		if haystack[i] == needle[j] {
			j++
		} else {
			for j > 0 && haystack[i] != needle[j] {
				j = next[j-1]
			}
			if haystack[i] == needle[j] {
				j++
			}
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func main() {
	s := []string{"aaaab", "", "ababa", "aabaaabaabaaab"}
	s2 := []string{"aab", "", "abb", "aabaaab"}
	for i := 0; i < len(s) && i < len(s2); i++ {
		fmt.Println("ret:", strStr(s[i], s2[i]))
	}
}
