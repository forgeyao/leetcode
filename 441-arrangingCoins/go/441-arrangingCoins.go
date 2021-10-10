// https://leetcode-cn.com/problems/arranging-coins/submissions/
package main

import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	return int((math.Sqrt(float64(1+8*n)) - 1) / 2)
}

func main() {
	n := []int{1, 5, 8, 1804289383}
	for _, i := range n {
		fmt.Println("ret:", arrangeCoins(i))
	}
}
