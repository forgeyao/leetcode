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
			ls := len(arr) - 1
			if ls >= 0 && (arr[ls] == '(' && s[i] == ')' || arr[ls] == '{' && s[i] == '}' || arr[ls] == '[' && s[i] == ']') {
				arr = arr[:ls]
			} else {
				return false
			}
		}
	}
	return len(arr) == 0
}

func main() {
	s := []string{"()", "()[]{}", "(]", "([)]", "{[]}", "(((", "(){}}{", "(]"}
	ans := []bool{true, true, false, false, true, false, false, false}
	for i := 0; i < len(s) && i < len(ans); i++ {
		r := isValid(s[i])
		fmt.Println(s[i], "ret:", r, ans[i])
	}
}
