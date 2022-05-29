// https://leetcode-cn.com/problems/happy-number/
package main

import "fmt"

func isHappy(n int) bool {
	record := map[int]bool{}
	//fmt.Print(n)
	for n != 1 {
		m := 0
		for i := n; i != 0; {
			tmp := i % 10
			m += tmp * tmp
			i = i / 10
		}
		//fmt.Print("->", m)
		if _, ok := record[m]; ok {
			//fmt.Println()
			return false
		}
		n = m
		record[m] = true
	}
	//fmt.Println()
	return n == 1
}

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 19}
	for _, val := range n {
		fmt.Println("ret:", isHappy(val))
	}
}
