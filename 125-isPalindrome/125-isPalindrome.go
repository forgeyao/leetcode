package main

import "fmt"

func isPalindrome(s string) bool {

	isAlphanumeric := func(c byte) bool {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			return true
		}
		return false
	}

	i, j := 0, len(s)-1
	for i < j {
		if !isAlphanumeric(s[i]) {
			i++
			continue
		}
		if !isAlphanumeric(s[j]) {
			j--
			continue
		}
		if s[i] == s[j] || (s[i] >= 'a' && s[i] <= 'z' && s[i] == s[j]+32) || (s[i] >= 'A' && s[i] <= 'Z' && s[i] == s[j]-32) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	a := []string{"A man, a plan, a canal: Panama", "race a car", "", " ", "  "}
	for i := 0; i < len(a); i++ {
		fmt.Println("ret:", isPalindrome(a[i]))
	}
}
