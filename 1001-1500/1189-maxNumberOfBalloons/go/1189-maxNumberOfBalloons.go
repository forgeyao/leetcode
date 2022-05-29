// https://leetcode-cn.com/problems/maximum-number-of-balloons/
package main

import "fmt"

/**
 * 对 text 里的字符计数，再看"balloon"里哪个字符数量最小
 * 时间O(n),n 为 text 长度
 * 空间O(C),C 为 text 字符种类数
 */
func maxNumberOfBalloons(text string) int {
	m := map[rune]int{}
	for _, ch := range text {
		m[ch]++
	}
	num := len(text)
	if m['b'] < num {
		num = m['b']
	}
	if m['a'] < num {
		num = m['a']
	}
	if m['l']/2 < num {
		num = m['l'] / 2
	}
	if m['o']/2 < num {
		num = m['o'] / 2
	}
	if m['n'] < num {
		num = m['n']
	}
	return num
}

/**
 * 官方答案，空间用得更少
 */
func maxNumberOfBalloons2(text string) int {
	m := [5]int{}
	for _, ch := range text {
		if ch == 'b' {
			m[0]++
		} else if ch == 'a' {
			m[1]++
		} else if ch == 'l' {
			m[2]++
		} else if ch == 'o' {
			m[3]++
		} else if ch == 'n' {
			m[4]++
		}
	}
	m[2] /= 2
	m[3] /= 2
	num := m[0]
	for _, v := range m {
		if v < num {
			num = v
		}
	}
	return num
}

func main() {
	texts := []string{"nlaebolko", "loonbalxballpoon", "leetcode"}
	for _, text := range texts {
		fmt.Println(text, " ", maxNumberOfBalloons(text), maxNumberOfBalloons2(text))
	}
}
