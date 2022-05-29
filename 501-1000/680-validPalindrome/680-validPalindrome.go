package main

import "fmt"

func validPalindrome(s string) bool {
	isPalindrome := func(sub string) bool {
		//fmt.Println("sub:", sub)
		for i, j := 0, len(sub)-1; i < j; i, j = i+1, j-1 {
			if sub[i] != sub[j] {
				return false
			}
		}
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return isPalindrome(s[i+1:j+1]) || isPalindrome(s[i:j])
		}
	}
	return true
}

func main() {
	a := []string{"aba", "abca", "", " ", "abcda", "abcdefcba"}

	for i := 0; i < len(a); i++ {
		fmt.Println("ret:", validPalindrome(a[i]))
	}
}
