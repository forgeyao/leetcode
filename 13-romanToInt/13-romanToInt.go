// https://leetcode-cn.com/problems/roman-to-integer/
package main

import "fmt"

func romanToInt(s string) int {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	num := m[s[0]]
	for i := 1; i < len(s); i++ {
		if m[s[i-1]] < m[s[i]] {
			num += m[s[i]] - m[s[i-1]] - m[s[i-1]]
		} else {
			num += m[s[i]]
		}
	}
	return num
}

func main() {
	s := []string{"III", "IV", "IX", "LVIII", "MCMXCIV"}
	for i := 0; i < len(s); i++ {
		fmt.Println("ret:", romanToInt(s[i]))
	}
}
