// https://leetcode-cn.com/problems/construct-the-rectangle/
package main

import (
	"fmt"
	"math"
)

// 借助开方,再迭代 l,w 到合适的值
func constructRectangle(area int) []int {
	sqrt := int(math.Sqrt(float64(area)))
	l, w := sqrt, sqrt
	for l*w != area && l > 0 && w > 0 {
		if l*w < area {
			l += 1
		} else {
			w -= 1
		}
	}
	return []int{l, w}
}

// 方法一 迭代效率低，改进版
func constructRectangle2(area int) []int {
	sqrt := int(math.Sqrt(float64(area)))
	l, w := sqrt, sqrt
	for l*w != area && l > 0 && w > 0 {
		if l*w < area {
			// 这里主要是会出现相等，意味着 l 已经到最大值了，需要调整 w
			if area/w <= l {
				w -= 1
			} else {
				l = area / w
			}
		} else {
			w = area / l
		}
		//fmt.Println("l,w: ", l, w)
	}
	return []int{l, w}
}

func main() {
	area := []int{4, 12, 15, 65, 77}
	for _, n := range area {
		fmt.Println("ret:", constructRectangle(n), constructRectangle2(n))
	}
}
