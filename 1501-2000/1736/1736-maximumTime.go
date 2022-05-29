// https://leetcode-cn.com/problems/latest-time-by-replacing-hidden-digits/
package main

import "fmt"

func maximumTime(time string) string {
	b := []byte(time)
	if b[0] == '?' {
		if b[1] >= '4' && b[1] <= '9' {
			b[0] = '1'
		} else {
			b[0] = '2'
		}
	}
	if b[1] == '?' {
		if b[0] == '2' {
			b[1] = '3'
		} else {
			b[1] = '9'
		}
	}
	if b[3] == '?' {
		b[3] = '5'
	}
	if b[4] == '?' {
		b[4] = '9'
	}
	return string(b)
}

func main() {
	s := []string{"2?:?0", "0?:3?", "1?:22", "??:3?", "??:??", "?4:03"}
	for _, v := range s {
		fmt.Println("ret:", maximumTime(v))
	}
}
