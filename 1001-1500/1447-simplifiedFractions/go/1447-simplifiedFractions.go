// https://leetcode-cn.com/problems/simplified-fractions/
package main

import (
	"fmt"
	"strconv"
)

/**
 * 题目难点在与要过于非最简分数，即分子与分母存在公约数
 * 最小公约数都是素数。考虑到 n<=100, 提前枚举出 n 内的素数
 * 时间O(C*n*n)
 * 空间 O(C), C 是 n 内的素数
 */
func simplifiedFractions(n int) []string {
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	ret := []string{}
	for i := 1; i < n; i++ { // 分子
		s := strconv.Itoa(i) + "/"
		for j := i + 1; j <= n; j++ { // 分母
			b := false
			for _, k := range prime {
				if k > i { // 公约数不会超过分子
					break
				}
				if i%k == 0 && j%k == 0 { // 存在公约数
					b = true
					break
				}
			}
			if b {
				continue
			}
			ret = append(ret, s+strconv.Itoa(j))
		}
	}
	return ret
}

/**
 * 官方答案
 * 利用最小公约数都是1，求解最简分数
 * 时间O(n*n*logn)
 * 空间O(1)
 */
func simplifiedFractions2(n int) (ans []string) {
	for denominator := 2; denominator <= n; denominator++ {
		for numerator := 1; numerator < denominator; numerator++ {
			if gcd(numerator, denominator) == 1 {
				ans = append(ans, strconv.Itoa(numerator)+"/"+strconv.Itoa(denominator))
			}
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func main() {
	ns := []int{2, 3, 4, 1, 22}
	for _, n := range ns {
		fmt.Println(n, ":", simplifiedFractions(n))
		fmt.Println(n, ":", simplifiedFractions2(n))
	}
}
