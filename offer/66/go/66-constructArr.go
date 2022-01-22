// https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/
package main

import "fmt"

func constructArr(a []int) []int {
	if len(a) <= 1 {
		return a
	}

	b := make([]int, len(a))
	b[0] = 1
	for i := 1; i < len(a); i++ {
		b[i] = b[i-1] * a[i-1]
	}

	for i, tmp := len(a)-2, 1; i >= 0; i-- {
		tmp *= a[i+1]
		b[i] *= tmp
	}
	return b
}

func main() {
	nums := [][]int{{1, 2, 3, 4, 5}, {1, 2}, {0}, {}}

	for _, num := range nums {
		fmt.Println("input:", num)
		fmt.Println("ret:", constructArr(num))
	}
}
