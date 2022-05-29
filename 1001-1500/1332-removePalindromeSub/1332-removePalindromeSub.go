package main

import "fmt"

func removePalindromeSubOld(s string) int {
	if s == "" {
		return 0
	}
	count := 0
	for mid := len(s) / 2; mid >= 0; mid-- {
		// odd
		left, right := 0, 0
		if len(s)%2 == 1 {
			left, right = mid-1, mid+1
		} else {
			left, right = mid-1, mid
		}
		//		fmt.Println("s,mid,l,r", s, mid, left, right)
		palindrom := false
		for ; left >= 0 && right < len(s); left, right = left-1, right+1 {
			if s[left] != s[right] {
				palindrom = false
				break
			}
			palindrom = true
		}
		//		fmt.Println("p", palindrom)
		if palindrom || mid == 0 {
			count++
			if 2*mid+1 >= len(s) {
				s = ""
				break
			} else {
				if mid == 0 || mid%2 == 1 {
					s = s[2*mid+1 : len(s)]
				} else {
					s = s[2*mid : len(s)]
				}
				mid = len(s)
			}
			continue
		}
	}
	return count
}

func removePalindromeSub(s string) int {
	if s == "" {
		return 0
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return 2
		}
	}
	return 1
}
func main() {
	s := []string{"ababa", "abb", "baabb", "", "ababb", "a", "ab", "aba", "abbaaaab"}
	for i := 0; i < len(s); i++ {
		fmt.Println("ret:", removePalindromeSub(s[i]))
	}
}
