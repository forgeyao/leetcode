//https://leetcode-cn.com/problems/sort-the-jumbled-numbers/
package main

import "sort"

func sortJumbled(mapping []int, nums []int) []int {
	mNums := []int{}
	mKv := map[int][]int{}
	for _, num := range nums {
		n, m, k := 0, 1, num
		if k == 0 {
			n = mapping[k]
		} else {
			for k != 0 {
				n += mapping[k%10] * m
				k /= 10
				m *= 10
			}
		}
		if v, ok := mKv[n]; ok {
			v = append(v, num)
			mKv[n] = v
		} else {
			arr := []int{num}
			mKv[n] = arr
		}
		mNums = append(mNums, n)
	}
	sort.Ints(mNums)
	//fmt.Println(mNums, mKv)
	ret := []int{}
	for i := 0; i < len(mNums); i++ {
		if i > 0 && mNums[i] == mNums[i-1] {
			continue
		}
		ret = append(ret, mKv[mNums[i]]...)
	}
	return ret
}
