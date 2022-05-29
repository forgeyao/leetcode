// https://leetcode-cn.com/problems/occurrences-after-bigram/
package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	str := strings.Split(text, " ")
	ret := []string{}
	for i := 2; i < len(str); i++ {
		if str[i-2] == first && str[i-1] == second {
			ret = append(ret, str[i])
		}
	}
	return ret
}

func main() {
	text := []string{"alice is a good girl she is a good student", "we will we will rock you"}
	first := []string{"a", "we"}
	second := []string{"good", "will"}

	for i := 0; i < len(text) && i < len(first) && i < len(second); i++ {
		fmt.Println("ret:", findOcurrences(text[i], first[i], second[i]))
	}
}
