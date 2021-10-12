// https://leetcode-cn.com/problems/truncate-sentence/
package main

import "fmt"

func truncateSentence(s string, k int) string {
	count := 0
	for i, c := range s {
		if c == ' ' {
			count++
			if count >= k {
				return s[:i]
			}
		}
	}
	return s
}

func main() {
	s := []string{"Hello how are you Contestant", "What is the solution to this problem"}
	k := []int{4, 4}
	for i := 0; i < len(s) && i < len(k); i++ {
		fmt.Println("ret:", truncateSentence(s[i], k[i]))
	}
}
