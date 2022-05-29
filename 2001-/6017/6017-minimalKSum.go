// https://leetcode-cn.com/problems/append-k-integers-with-minimal-sum/
package main

import (
	"fmt"
	"sort"
)

/**
 * 先用排序,再根据排序补充数据
 */
func minimalKSum(nums []int, k int) int64 {
	sort.Ints(nums)
	fmt.Println(nums)
	var ret int64 = 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if k < nums[i] {
				return int64((1 + k) * k / 2)
			}
			ret = int64(nums[i] * (nums[i] - 1) / 2)
			k -= (nums[i] - 1)
		} else {
			if nums[i]-nums[i-1] <= 1 {
				continue
			}
			if k < nums[i]-nums[i-1]-1 {
				ret += int64((nums[i-1]*2 + k + 1) * k / 2)
				k = 0
			} else {
				ret += int64((nums[i-1] + 1 + nums[i] - 1) * (nums[i] - nums[i-1] - 1) / 2)
				k -= nums[i] - nums[i-1] - 1
			}
		}
		fmt.Println(ret)
		if k == 0 {
			break
		}
	}
	if k != 0 {
		last := nums[len(nums)-1]
		ret += int64((last + 1 + last + k) * k / 2)
	}
	return ret

	/*
		 * 没用排序，超时了
		 *
		 m := map[int64]int{}
		  min := nums[0]
		  for _, num := range nums {
			  m[int64(num)]++
			  if num < min {
				  min = num
			  }
		  }
		  var ret int64 = 0
		  if k < min {
			  return int64((1+k)*k/2)
		  } else {
			  ret = int64(min*(min-1)/2)
			  k -= (min-1)
		  }
		  var i int64 = int64(min+1)
		  for ; k > 0; i++ {
			  if _,ok := m[i]; ok {
				  continue
			  } else {
				  ret += i
				  k--
			  }
		  }
		  return ret*/
}
