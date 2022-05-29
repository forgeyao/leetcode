// https://leetcode-cn.com/problems/longest-common-prefix/
package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	min := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if min > len(strs[i]) {
			min = len(strs[i])
		}
	}

	for i := 0; i < min; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return strs[0][0:i]
			}
		}
	}
	return strs[0][0:min]
}

func main() {
	s := [][]string{{"flower", "flow", "flight"}, {"dog", "racecar", "car"}, {"abd", "abd", "abd"}}
	for i := 0; i < len(s); i++ {
		fmt.Println("ret:", longestCommonPrefix(s[i]))
	}
}
