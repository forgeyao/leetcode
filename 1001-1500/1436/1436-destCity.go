// https://leetcode-cn.com/problems/destination-city/
package main

import "fmt"

func destCity(paths [][]string) string {
	m := make(map[string]bool, len(paths))
	for _, v := range paths {
		m[v[0]] = true
	}
	for _, v := range paths {
		if _, ok := m[v[1]]; !ok {
			return v[1]
		}
	}
	return ""
}

func main() {
	s := [][][]string{{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}, {{"B", "C"}, {"D", "B"}, {"C", "A"}}, {{"A", "Z"}}}
	for _, v := range s {
		fmt.Println("ret:", destCity(v))
	}
}
