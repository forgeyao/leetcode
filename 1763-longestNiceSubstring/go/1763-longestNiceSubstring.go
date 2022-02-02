// https://leetcode-cn.com/problems/longest-nice-substring/
package main

import (
	"fmt"
	"unicode"
)

/**
 * 常规思维
 */
func longestNiceSubstring(s string) (ans string) {
	for i := range s {
		lower, upper := 0, 0
		for j := i; j < len(s); j++ {
			if unicode.IsLower(rune(s[j])) {
				lower |= 1 << (s[j] - 'a')
			} else {
				upper |= 1 << (s[j] - 'A')
			}
			if lower == upper && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return
}

/**
 * 分治法
 */
func longestNiceSubstring2(s string) (ans string) {
	if s == "" {
		return
	}
	lower, upper := 0, 0
	for _, ch := range s {
		if unicode.IsLower(ch) {
			lower |= 1 << (ch - 'a')
		} else {
			upper |= 1 << (ch - 'A')
		}
	}
	if lower == upper {
		return s
	}
	valid := lower & upper
	for i := 0; i < len(s); i++ {
		start := i
		for i < len(s) && valid>>(unicode.ToLower(rune(s[i]))-'a')&1 == 1 {
			i++
		}
		if t := longestNiceSubstring2(s[start:i]); len(t) > len(ans) {
			ans = t
		}
	}
	return
}

func main() {
	ss := []string{"YazaAay", "Bb", "c", "dDzeE"}
	for _, s := range ss {
		fmt.Println(s, " ret:", longestNiceSubstring(s), longestNiceSubstring2(s))
	}
}
