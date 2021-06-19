// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
package main

import "fmt"

func maxProfit2(prices []int) int {
	length := len(prices)
	min, max, diff := 0, 0, 0
	for i := 1; i < length; {
		if prices[min] >= prices[i] {
			min = i
			i++
			continue
		} else {
			max = i
			for j := i + 1; j < length; j++ {
				if prices[j] > prices[max] {
					max = j
				}
			}
			for j := i + 1; j < max; j++ {
				if prices[min] > prices[j] {
					min = j
				}
			}
		}
		if max > min && prices[max]-prices[min] > diff {
			diff = prices[max] - prices[min]
		}
		i = max + 2
		min = max + 1
		max = max + 1
	}
	return diff
}

func maxProfit(prices []int) int {
	min, maxDiff := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxDiff {
			maxDiff = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxDiff
}

func main() {
	n := [][]int{{7, 1, 5, 3, 6, 4}, {7, 6, 4, 3, 1}, {1, 2, 3, 4, 5}, {4, 5, 6, 7, 1, 5}, {4, 5, 2, 7, 8}}
	for i := 0; i < len(n); i++ {
		fmt.Println("ret:", maxProfit(n[i]))
	}
}
