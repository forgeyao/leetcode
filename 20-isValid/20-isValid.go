// https://leetcode-cn.com/problems/valid-parentheses/
package main

import "fmt"

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	arr := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			arr = append(arr, s[i])
		} else {
			if len(arr) < 1 {
				return false
			}
			c := arr[len(arr)-1]
			if c == '(' && s[i] == ')' || c == '{' && s[i] == '}' || c == '[' && s[i] == ']' {
				arr = arr[:len(arr)-1]
				continue
			}
			return false
		}
	}
	return len(arr) == 0
}

func main() {
	s := []string{"()", "()[]{}", "(]", "([)]", "{[]}", "((("}
	for i := 0; i < len(s); i++ {
		fmt.Println("ret:", isValid(s[i]))
	}
}
