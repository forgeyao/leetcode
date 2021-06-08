// https://leetcode-cn.com/problems/add-binary/
package main

import "fmt"

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	max := m
	if max < n {
		max = n
	}
	ret := make([]byte, max+1) // 多申请一位

	var flag byte
	for i, j, k := m-1, n-1, max; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		var an, bn byte = 0, 0
		if i >= 0 {
			an = a[i] - '0'
		}
		if j >= 0 {
			bn = b[j] - '0'
		}

		if an+bn+flag <= 1 {
			ret[k] = an + bn + flag + '0'
			flag = 0
		} else {
			ret[k] = an + bn + flag + '0' - 2
			flag = 1
		}
	}
	if flag == 1 {
		ret[0] = '1'
	} else {
		ret = ret[1:]
	}
	return string(ret)
}

func main() {
	s1 := []string{"11", "1010", "1"}
	s2 := []string{"1", "1011", "111"}
	for i := 0; i < len(s1) && i < len(s2); i++ {
		fmt.Println("ret:", addBinary(s1[i], s2[i]))
	}
}
