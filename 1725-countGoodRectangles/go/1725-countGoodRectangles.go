// https://leetcode-cn.com/problems/number-of-rectangles-that-can-form-the-largest-square/
package main

import "fmt"

/**
 * 遍历矩形数组，获取较短的一边，再在短边数组里对最大值计数
 * 时间 O(n)
 * 空间 O(1)
 */
func countGoodRectangles(rectangles [][]int) int {
	max, count := 0, 0
	for _, rect := range rectangles {
		if rect[0] > rect[1] {
			rect[0], rect[1] = rect[1], rect[0]
		}
		if max < rect[0] {
			max = rect[0]
			count = 1
		} else if max == rect[0] {
			count++
		}
	}
	return count
}

func main() {
	rectangles := [][][]int{
		{{5, 8}, {3, 9}, {5, 12}, {16, 5}},
		{{2, 3}, {3, 7}, {4, 3}, {3, 7}},
	}
	for _, rect := range rectangles {
		fmt.Println(countGoodRectangles(rect))
	}
}
