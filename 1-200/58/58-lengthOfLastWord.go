// https://leetcode-cn.com/problems/length-of-last-word/
package main

import "fmt"

func lengthOfLastWord(s string) int {
	idx := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if idx == -1 {
				continue
			}
			return idx - i
		}
		if idx == -1 {
			idx = i
		}
	}
	return idx + 1
}

func main() {
	s := []string{"Hello World", " ", "  abc ", "abc    "}
	for i := 0; i < len(s); i++ {
		fmt.Println("ret:", lengthOfLastWord(s[i]))
	}
}
