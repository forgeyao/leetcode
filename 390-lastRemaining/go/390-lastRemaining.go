// https://leetcode-cn.com/problems/elimination-game/
package main

import "fmt"

/**
 * 看起来很简单，实现起来挺难的
 * 没有特别好的思路，下面是直接抄答案的
 */
func lastRemaining(n int) int {
	a1 := 1                 // 数组首个元素
	k, cnt, step := 0, n, 1 // k:第k次删除  cnt:元素个数  step:等差数列的公差
	for cnt > 1 {
		if k%2 == 0 {
			a1 += step
		} else {
			if cnt%2 == 1 {
				a1 += step
			}
		}
		k++
		cnt >>= 1
		step <<= 1
	}
	return a1
}

func main() {
	for i := 0; i < 40; i++ {
		fmt.Println(i, " ret:", lastRemaining(i))
	}
}
