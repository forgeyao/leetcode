// https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/
package main

import (
	"fmt"
)

/**
 * 思路：为了控制索引和最小，从小到大遍历索引和
 * 时间 O((l1+l2)*l1*r) l1/l2 表示list1/list2 长度,x 表示字符串平均长度
 * 空间 O(1)
 */
func findRestaurant(list1 []string, list2 []string) []string {
	ret := []string{}
	for sum := 0; sum <= len(list1)+len(list2)-2; sum++ { // sum 表示索引和，从最小索引开始向上遍历
		for i, j := 0, sum; i < len(list1) && j >= 0; {
			if j < len(list2) && list1[i] == list2[j] {
				ret = append(ret, list1[i])
			}
			i++
			j = sum - i
		}
		if len(ret) > 0 {
			break
		}
	}
	return ret
}

/**
 * 官方解答一，利用哈希表缓存
 * 时间 O(l1 ∗l2 ∗x)。l1和 l2 是 list1 和 list2的长度，x 是字符串的平均长度。
 * 空间 O(l1 *l2 *x)
 */
func findRestaurant2(list1 []string, list2 []string) []string {
	m := map[int][]string{}
	for i := 0; i < len(list1); i++ {
		for j := 0; j < len(list2); j++ {
			if list1[i] == list2[j] {
				m[i+j] = append(m[i+j], list1[i])
			}
		}
	}
	min := len(list1) + len(list2) - 2
	for k := range m {
		if k < min {
			min = k
		}
	}
	return m[min]
}

/**
 * 官方解答三，利用哈希表缓存
 * 时间 O(l1 + l2)。l1和 l2 是 list1 和 list2的长度，x 是字符串的平均长度。
 * 空间 O(l1 *x)
 */
func findRestaurant3(list1 []string, list2 []string) []string {
	m := map[string]int{}
	for i := 0; i < len(list1); i++ {
		m[list1[i]] = i
	}
	ret := []string{}
	min := len(list1) + len(list2) - 2
	for j := 0; j < len(list2); j++ {
		if k, ok := m[list2[j]]; ok {
			if k+j < min {
				ret = ret[:0]
				ret = append(ret, list2[j])
				min = k + j
			} else if k+j == min {
				ret = append(ret, list2[j])
			}
		}
	}
	return ret
}

func main() {
	list1 := [][]string{
		{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		{"Shogun", "Tapioca Express", "Burger King", "KFC"},
		{"vacag", "KFC"},
		{"Shogun", "Piatti", "Tapioca Express", "Burger King", "KFC"},
	}
	list2 := [][]string{
		{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
		{"KFC", "Shogun", "Burger King"},
		{"KNN", "KFC", "Burger King", "Tapioca Express", "Shogun"},
		{"fvo", "xrljq", "jrl", "KFC"},
		{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"},
	}
	for i := 0; i < len(list1) && i < len(list2); i++ {
		fmt.Println("ret:", findRestaurant(list1[i], list2[i]), findRestaurant2(list1[i], list2[i]), findRestaurant3(list1[i], list2[i]))
	}
}
