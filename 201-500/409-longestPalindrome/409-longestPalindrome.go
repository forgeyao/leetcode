package main

import "fmt"

func longestPalindrome(s string) int {
	var a [52]int
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			a[s[i]-'A'] += 1
		} else {
			a[s[i]-71] += 1
		}
	}
	sum := 0
	for i := 0; i < 52; i++ {
		sum += (a[i] / 2) * 2
	}
	if len(s) > sum {
		sum += 1
	}
	return sum
}

func main() {
	s := []string{"", "abcd", "abccccdd", "bacdefadkfjadfddddABACDAzz"}
	for i := 0; i < len(s); i++ {
		fmt.Println("ret:", longestPalindrome(s[i]))
	}
}
