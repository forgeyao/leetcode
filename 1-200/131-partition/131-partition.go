/**
* https://leetcode-cn.com/problems/palindrome-partitioning/
 */
package main

import "fmt"

func partition(s string) [][]string {
	isPalindrome := func(sub string) bool {
		for i, j := 0, len(sub)-1; i < j; i, j = i+1, j-1 {
			if sub[i] != sub[j] {
				return false
			}
		}
		return true
	}

	var part func(s string) [][]string
	part = func(s string) [][]string {
		var res [][]string
		if isPalindrome(s) {
			ss := []string{s}
			res = append(res, ss)
		}
		if len(s) == 1 {
			return res
		}

		for i := 1; i < len(s); i++ {
			left, right := s[0:i], s[i:len(s)]

			if !isPalindrome(left) {
				continue
			}
			//fmt.Println("left:", left)
			rightSub := part(right)
			//fmt.Println("right:", rightSub)
			for k := 0; k < len(rightSub); k++ {
				ss := []string{left}
				ss = append(ss, rightSub[k]...)
				res = append(res, ss)
			}
		}
		return res
	}

	var res [][]string
	res = part(s)
	return res
}

func main() {
	s := []string{"a", "aa", "ab", "aab", "aaaaa", "aababa"}
	for i := 0; i < len(s); i++ {
		fmt.Println("ret:", partition(s[i]))
	}
}
