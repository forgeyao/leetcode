package main

import "fmt"

//brute force
func longestPalindromeOld(s string) string {
	isPalindrome := func(sub string) bool {
		for i, j := 0, len(sub)-1; i < j; i, j = i+1, j-1 {
			if sub[i] != sub[j] {
				return false
			}
		}
		return true
	}

	for i := len(s); i >= 0; i-- {
		beg, end := 0, i
		for ; end <= len(s); beg, end = beg+1, end+1 {
			if isPalindrome(s[beg:end]) {
				return s[beg:end]
			}
		}
	}
	return ""
}

// Center extension
func longestPalindrome2(s string) string {
	i, j := 0, 0
	for mid := 1; mid < len(s); mid++ {
		//odd
		left, right := mid-1, mid+1
		for ; left >= 0 && right < len(s); left, right = left-1, right+1 {
			if s[left] != s[right] {
				break
			}
			if right-left > j-i {
				i, j = left, right
			}
		}
		// even
		left, right = mid-1, mid
		for ; left >= 0 && right < len(s); left, right = left-1, right+1 {
			if s[left] != s[right] {
				break
			}
			if right-left > j-i {
				i, j = left, right
			}
		}
	}
	return s[i : j+1]
}

// Dynamic programming
func longestPalindrome(s string) string {
	a := make([][]bool, len(s))

	for i := 0; i < len(s); i++ {
		a[i] = make([]bool, len(s))
		a[i][i] = true
	}

	left, right := 0, 0
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if j-i == 1 {
				if s[j] == s[i] {
					a[i][j] = true
				} else {
					a[i][j] = false
				}
			} else {
				a[i][j] = (a[i+1][j-1] && (s[i] == s[j]))
			}
			if a[i][j] && (j-i > right-left) {
				left, right = i, j
			}
		}
	}
	return s[left : right+1]
}

func main() {
	a := []string{"babad", "cbbd", "a", "ac", "bb"}
	for i := 0; i < len(a); i++ {
		fmt.Println("ret:", longestPalindrome(a[i]))
	}
}
