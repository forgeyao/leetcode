// https://leetcode-cn.com/problems/minimum-number-of-steps-to-make-two-strings-anagram-ii/
package main

import "fmt"

/**
 * 思路：用哈希缓存字符数量，再对比两个哈希表差异
 * 时间 O(M+N+m+n) M,N 表示 s,t 字符串长度 m,n 表示 s,t 字符种类数
 * 空间 O(m+n)
 */
func minSteps(s string, t string) int {
	ms := map[rune]int{}
	mt := map[rune]int{}
	for _, ch := range s {
		ms[ch]++
	}
	for _, ch := range t {
		mt[ch]++
	}
	ret := 0
	for k, v := range ms {
		if vv, ok := mt[k]; ok {
			if v > vv {
				ret += v - vv
			} else {
				ret += vv - v
			}
		} else {
			ret += v
		}
	}
	for k, v := range mt {
		if _, ok := ms[k]; !ok {
			ret += v
		}
	}
	return ret
}

func main() {
	ss := []string{"leetcode", "night"}
	ts := []string{"coats", "thing"}
	for i := 0; i < len(ss) && i < len(ts); i++ {
		fmt.Println(minSteps(ss[i], ts[i]))
	}
}
