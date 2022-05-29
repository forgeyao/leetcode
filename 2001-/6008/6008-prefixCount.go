// https://leetcode-cn.com/problems/counting-words-with-a-given-prefix/
package main

import (
	"fmt"
	"strings"
)

/**
 * 时间 O(m*n), m 是 words 长度, n 是 pref 长度
 * 空间 O(1)
 */
func prefixCount(words []string, pref string) int {
	count := 0
	for i := 0; i < len(words); i++ {
		//if len(words[i]) >= l && string(words[i][:l]) == pref {
		if strings.HasPrefix(words[i], pref) {
			count++
		}
	}
	return count
}

func main() {
	words := [][]string{
		{"pay", "attention", "practice", "attend"},
		{"leetcode", "win", "loops", "success"},
	}
	prefs := []string{"at", "code"}
	for i := 0; i < len(words) && i < len(prefs); i++ {
		fmt.Println(prefixCount(words[i], prefs[i]))
	}
}
