// https://leetcode-cn.com/problems/find-missing-observations/
package main

import "fmt"

/**
 * 将题简化为对整数sum, 分解为n 个数, 数的范围是[1,6]
 *
 * 思路1
 * 递归。将求解n个数，转换为求 n-1 和 1个数的情况
 * 这种解法效率不够，没通过最长的那个 test case
 * 中间有很多重复计算，可用动态规划优化，但我直接放弃这种方案
 *
 * 思路2
 * 平均数。avg = sum/n (avg 向下取整)
 * 要求解 n 个数，可以直接构造 n 个 avg，再补充 sum - n*avg 个数就行
 */
func missingRolls(rolls []int, mean int, n int) []int {
	sum := (len(rolls) + n) * mean
	for i := range rolls {
		sum -= rolls[i]
	}
	if sum < n || sum > 6*n {
		return []int{}
	}

	var avg int
	avg = sum / n
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = avg
	}

	left := sum - avg*n
	diff := 6 - avg
	for j := 0; j < n && left > 0 && diff > 0; j++ {
		if left >= diff {
			ret[j] += diff
			left -= diff
		} else {
			ret[j] += left
			left = 0
		}
	}
	if left > 0 {
		return []int{}
	}
	return ret
}

func getRolls(sum int, n int) []int {
	if n == 1 {
		if sum < 1 || sum > 6 {
			return []int{}
		}
		return []int{sum}
	}

	for i := 1; i <= 6; i++ {
		ret := getRolls(i, 1)
		ret = append(ret, getRolls(sum-i, n-1)...)
		//		fmt.Println("getRolls, i:", i, ret)
		if len(ret) == n {
			return ret
		}
	}
	return []int{}
}

func main() {
	rolls := [][]int{{3, 2, 4, 3}, {1, 5, 6}, {1, 2, 3, 4}, {1}, {4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}}
	mean := []int{4, 3, 6, 3, 4}
	n := []int{2, 4, 4, 1, 40}
	for i := 0; i < len(rolls) && i < len(mean) && i < len(n); i++ {
		fmt.Println("rolls,mean,n:", rolls[i], mean[i], n[i], " ret:", missingRolls(rolls[i], mean[i], n[i]))
	}
}
