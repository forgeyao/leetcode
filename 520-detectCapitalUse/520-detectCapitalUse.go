// https://leetcode-cn.com/problems/detect-capital/
package main

import "fmt"

/**
 * word 里有 n 个元素
 * [1, n) 之间要么全是大写要么全是小写才为 true，否则为 false
 * 最后还有排查一种情况，首字母是小写后续都是大写的情况 如"aBC"
 */
func isBig(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

func detectCapitalUse(word string) bool {
	l := len(word)
	last := isBig(word[l-1])
	for i := l - 2; i >= 1; i-- {
		if last != isBig(word[i]) {
			return false
		}
	}
	if last && !isBig(word[0]) {
		return false
	}

	return true
}

func main() {
	words := []string{"UAS", "FlaG", "leetcode", "aBC", "a", "A"}
	for _, word := range words {
		fmt.Println(word, detectCapitalUse(word))
	}
}
