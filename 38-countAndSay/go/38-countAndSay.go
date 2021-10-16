package main

import (
	"fmt"
	"strconv"
)

// 递归
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	s := countAndSay(n - 1)

	var ret string
	ch, count := s[0], 1
	for i := 1; i < len(s); i++ {
		if s[i] == ch {
			count++
			continue
		}

		ret += strconv.Itoa(count) + string(ch)
		count = 1
		ch = s[i]
	}

	return ret + strconv.Itoa(count) + string(ch)
}

// 非递归
func countAndSay2(n int) string {
	ret := "1"
	if n == 1 {
		return ret
	}

	countStr := func(s string) string {
		var ret string
		ch, count := s[0], 1
		for i := 1; i < len(s); i++ {
			if s[i] == ch {
				count++
				continue
			}

			ret += strconv.Itoa(count) + string(ch)
			count = 1
			ch = s[i]
		}
		return ret + strconv.Itoa(count) + string(ch)
	}

	for i := 2; i <= n; i++ {
		ret = countStr(ret)
	}

	return ret
}

// 题目 n 范围是  1 ～ 30
// 所以可以提前算出来存到 map。不过这出题意图，就不实现了

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("ret:", i, ",", countAndSay(i), countAndSay2(i))
	}
}
