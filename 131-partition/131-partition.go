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

	part := func(s string) [][]string {
		res := make([][]string, 16)
		if isPalindrome(s) {
			ss := make([]string, 1)
			ss = append(ss, s)
			res = append(res, ss)
		}
		if len(s) == 1 {
			return res
		}

		sub := part(s[1:len(s)])
		for i := 0; i < len(sub); i++ {
			sub[i] = append([]string{string(s[0])}, sub[i])
		}
		return append(res, sub)
	}

	res := make([][]string, 16)
	for i := 0; i < len(s); i++ {
		left, right := "", ""
		if i == 0 {
			left = ""
			right = s
		} else {
			left = s[0:i]
			right = s[i:len(s)]
		}

		leftSub := part(left)
		rightSub := part(right)
		for j := 0; j < len(leftSub); j++ {
			for k := 0; k < len(rightSub); k++ {
				res = append(res, leftSub[j], rightSub[k])
			}
		}
	}
	return res
}

func main() {
	s := []string{"aab", "a"}
	for i := 0; i < len(s); i++ {
		fmt.Println("ret:", partition(s[i]))
	}
}
