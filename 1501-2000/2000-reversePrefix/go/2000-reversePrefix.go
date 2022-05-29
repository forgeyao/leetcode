// https://leetcode-cn.com/problems/reverse-prefix-of-word/
package main

import "fmt"

func reversePrefix(word string, ch byte) string {
	i := -1
	for j := 0; j < len(word); j++ {
		if byte(word[j]) == ch {
			i = j
			break
		}
	}

	var ret []byte
	for j := i; j >= 0; j-- {
		ret = append(ret, word[j])
	}
	ret = append(ret, word[i+1:]...)
	return string(ret)
}

func main() {
	words := []string{"abcdefd", "xyxzxe", "abcd"}
	chs := []byte{'d', 'z', 'z'}
	for i := 0; i < len(words) && i < len(chs); i++ {
		fmt.Println(words[i], chs[i], " ret:", reversePrefix(words[i], chs[i]))
	}
}
