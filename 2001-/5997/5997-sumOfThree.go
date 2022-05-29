// https://leetcode-cn.com/problems/find-three-consecutive-integers-that-sum-to-a-given-number/
package main

import "fmt"

/**
 * 没想明白这题为啥是中等难度
 */
func sumOfThree(num int64) []int64 {
	ret := []int64{}
	if num%3 == 0 {
		avg := num / 3
		ret = append(ret, avg-1, avg, avg+1)
	}
	return ret
}

func main() {
	nums := []int64{33, 4}
	for _, num := range nums {
		fmt.Println(sumOfThree(num))
	}
}
