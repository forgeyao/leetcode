// https://leetcode-cn.com/problems/palindromic-substrings/
package main

import "fmt"

func countSubstrings(s string) int {
	count := 0
	l := len(s)
	cond := make([][]bool, l)
	for i := 0; i < l; i++ {
		cond[i] = make([]bool, l)
		cond[i][i] = true
		count++
	}
	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			if i+1 > j-1 {
				cond[i][j] = (s[i] == s[j])
			} else {
				cond[i][j] = (cond[i+1][j-1] && (s[i] == s[j]))
			}
			if cond[i][j] {
				count++
			}
		}
	}
	return count
}

func main() {
	s := []string{"abc", "aaa", "", "a", "aa"}
	for i := 0; i < len(s); i++ {
		fmt.Println("ret:", countSubstrings(s[i]))
	}
}
