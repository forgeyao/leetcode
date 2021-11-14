// https://leetcode-cn.com/problems/map-sum-pairs/
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MapSum struct {
	m map[string]int
}

func Constructor() MapSum {
	return MapSum{map[string]int{}}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	sum := 0
	for k, v := range this.m {
		//if len(k) >= lenPre && k[0:lenPre] == prefix {
		if strings.HasPrefix(k, prefix) {
			sum += v
		}
	}
	return sum
}

func main() {
	//command := []string{"MapSum", "insert", "sum", "insert", "sum", "sum"}
	//input := [][]string{{}, {"apple", "3"}, {"ap"}, {"app", "2"}, {"ap"}, {"a"}}

	command := []string{"MapSum", "insert", "sum", "insert", "sum", "sum", "insert", "sum", "sum", "sum", "insert", "sum", "sum", "sum", "sum", "sum", "insert", "insert", "insert", "sum", "sum", "sum", "sum", "sum", "sum", "insert", "sum", "sum"}
	input := [][]string{{}, {"aa", "3"}, {"a"}, {"aa", "2"}, {"a"}, {"aa"}, {"aaa", "3"}, {"aaa"}, {"bbb"}, {"ccc"}, {"aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj", "111"}, {"aa"}, {"a"}, {"b"}, {"c"}, {"aewfwaefjeoawefjwoeajfowajfoewajfoawefjeowajfowaj"}, {"bb", "143"}, {"cc", "144"}, {"aew", "145"}, {"bb"}, {"cc"}, {"aew"}, {"dddd"}, {"cdcd"}, {"aab"}, {"aab", "33"}, {"aab"}, {"ab"}}

	var mapSum MapSum
	for i := 0; i < len(command) && i < len(input); i++ {
		cmd := command[i]
		if cmd == "MapSum" {
			mapSum = Constructor()
		} else if cmd == "insert" {
			if len(input[i]) < 2 {
				fmt.Println("input insert size error")
				break
			}
			val, _ := strconv.Atoi(input[i][1])
			mapSum.Insert(input[i][0], val)
		} else if cmd == "sum" {
			if len(input[i]) < 1 {
				fmt.Println("input sum size error")
				break
			}
			fmt.Println("sum:", mapSum.Sum(input[i][0]))
		}
	}
}
