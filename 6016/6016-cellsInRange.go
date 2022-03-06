// https://leetcode-cn.com/problems/cells-in-a-range-on-an-excel-sheet/
package main

func cellsInRange(s string) []string {
	ret := []string{}
	for c := s[0]; c <= s[3]; c++ {
		tmp := []byte{c}
		for r := s[1]; r <= s[4]; r++ {
			if len(tmp) == 1 {
				tmp = append(tmp, r)
			} else {
				tmp[1] = r
			}
			ret = append(ret, string(tmp))
		}
	}
	return ret
}
