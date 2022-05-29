// https://leetcode-cn.com/problems/decompress-run-length-encoded-list/
package main

import "fmt"

func decompressRLElist(nums []int) []int {
	var ret []int
	for i := 0; i+1 < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			ret = append(ret, nums[i+1])
		}
	}
	return ret
}

func main() {
	nums := [][]int{{1, 2, 3, 4}, {1, 1, 2, 3}}
	for _, num := range nums {
		fmt.Println(decompressRLElist(num))
	}
}
