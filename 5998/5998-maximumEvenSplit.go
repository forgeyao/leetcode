// https://leetcode-cn.com/problems/maximum-split-of-positive-even-integers/
package main

import "fmt"

/**
 * 思路：要使数目最多，那就尽可能多使用小点的偶数
 * 故从2开始向上迭代
 */
func maximumEvenSplit(finalSum int64) []int64 {
	ret := []int64{}
	if finalSum%2 == 1 {
		return ret
	}
	var inc int64 = 2
	for finalSum != 0 {
		if finalSum <= 2*inc {
			ret = append(ret, finalSum)
			break
		}
		finalSum -= inc
		ret = append(ret, inc)
		inc += 2
	}
	return ret
}

func main() {
	finalSums := []int64{12, 7, 28}
	for _, finalSum := range finalSums {
		fmt.Println(maximumEvenSplit(finalSum))
	}
}
