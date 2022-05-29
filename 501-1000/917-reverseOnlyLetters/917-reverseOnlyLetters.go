// https://leetcode-cn.com/problems/reverse-only-letters/
package main

import "fmt"

/**
 * 双指针遍历一遍
 * 时间O(n)
 * 空间O(n), 不能直接修改原字符串
 */
func reverseOnlyLetters(s string) string {
	ret := make([]byte, len(s))
	for i, j := 0, len(s)-1; i <= j; {
		if !(s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z') {
			ret[i] = s[i]
			i++
			continue
		}
		if !(s[j] >= 'a' && s[j] <= 'z' || s[j] >= 'A' && s[j] <= 'Z') {
			ret[j] = s[j]
			j--
			continue
		}
		ret[i], ret[j] = s[j], s[i]
		i++
		j--
	}
	return string(ret)
}

func main() {
	ss := []string{"ab-cd", "a-bC-dEf-ghIj", "Test1ng-Leet=code-Q!"}
	for _, s := range ss {
		fmt.Println(reverseOnlyLetters(s))
	}
}
