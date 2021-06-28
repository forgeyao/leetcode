// https://leetcode-cn.com/problems/excel-sheet-column-title/
package main

import "fmt"

func convertToTitle(columnNumber int) string {
	m := map[int]byte{1: 'A', 2: 'B', 3: 'C', 4: 'D', 5: 'E', 6: 'F', 7: 'G', 8: 'H', 9: 'I', 10: 'J', 11: 'K', 12: 'L', 13: 'M', 14: 'N', 15: 'O', 16: 'P', 17: 'Q', 18: 'R', 19: 'S', 20: 'T', 21: 'U', 22: 'V', 23: 'W', 24: 'X', 25: 'Y', 0: 'Z'}

	var s string
	for columnNumber > 0 {
		remainder := columnNumber % 26
		s = string(m[remainder]) + s
		columnNumber = columnNumber / 26
		// 余数为0并不是真的0,而是26,所以要减掉 26
		if remainder == 0 {
			columnNumber -= 1
		}
	}
	return s
}

func main() {
	n := []int{1, 26, 28, 52, 701, 1000}
	for _, val := range n {
		fmt.Println("ret:", convertToTitle(val))
	}
}
