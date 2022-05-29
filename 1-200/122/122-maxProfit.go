// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
package main

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	min, length := prices[0], len(prices)
	for i := 1; i < length; i++ {
		if prices[i] < prices[i-1] {
			profit += prices[i-1] - min
			min = prices[i]
		}
	}
	if min < prices[length-1] {
		profit += prices[length-1] - min
	}
	return profit
}

func main() {
	n := [][]int{{7, 1, 5, 3, 6, 4}, {1, 2, 3, 4, 5}, {7, 6, 4, 3, 1}}
	for i := 0; i < len(n); i++ {
		fmt.Println("ret:", maxProfit(n[i]))
	}
}
