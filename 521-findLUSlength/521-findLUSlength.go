// https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/
package main

import "fmt"

/**
 * 重点在理解题意
 * 时间 O(n), n 是a、b 相同时字符串长度
 * 空间 O(1)
 */
func findLUSlength(a string, b string) int {
	la := len(a)
	lb := len(b)
	if la > lb {
		return la
	} else if lb > la {
		return lb
	}
	if a == b {
		return -1
	}
	return la
}

func main() {
	params := []struct {
		a      string
		b      string
		result int
	}{
		{"aba", "cdc", 3},
		{"aaa", "bbb", 3},
		{"aaa", "aaa", -1},
		{"horbxeemlgqpqbujbdagizcfairalg", "iwvtgyojrfhyzgyzeikqagpfjoaeen", 30},
	}
	for _, item := range params {
		ret := findLUSlength(item.a, item.b)
		fmt.Println(item.a, item.b, "result:", item.result, "ret:", ret)
	}
}
