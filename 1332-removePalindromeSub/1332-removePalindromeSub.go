package main

import "fmt"

func removePalindromeSub(s string) int {
	isPalindrome := func(sub string) bool {
		for i, j := 0, len(sub)-1; i < j; i, j = i+1, j-1 {
			if sub[i] != sub[j] {
				return false
			}
		}
		return true
	}

	count := 0
	for s != "" {
		i, j := 0, len(s)-1
		for i < j {
			if s[i] == s[j] {
				if isPalindrome(s[i+1 : j]) {
					break
				}
			}

			j = j - 1
			if i < j && s[i] == s[j] {
				if isPalindrome(s[i+1 : j]) {
					break
				}
			}

			i = i + 1
			j = j + 1
			if i < j && s[i] == s[j] {
				if isPalindrome(s[i+1 : j]) {
					break
				}
			}
			j = j - 1
		}
		//fmt.Println("i,j:", i, j)

		if i < j {
			s = s[0:i] + s[j:len(s)-1]
			count++
		}
		if s == "a" || s == "b" {
			s = ""
			count++
		} else if s == "ab" || s == "ba" {
			s = ""
			count += 2
		}
	}
	return count
}

func main() {
	s := []string{"ababa", "abb", "baabb", "", "ababb", "a", "ab", "aba"}
	for i := 0; i < len(s); i++ {
		fmt.Println("ret:", removePalindromeSub(s[i]))
	}
}
