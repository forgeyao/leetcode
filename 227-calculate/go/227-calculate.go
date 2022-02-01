// https://leetcode-cn.com/problems/basic-calculator-ii/
package main

import (
	"fmt"
	"strconv"
)

/**
 * 总体思路与官方解答一致, 都是用栈缓存数据
 * 但在细节处理上复杂很多
 */
func calculate(s string) int {
	var nums []int   // 数字栈
	var opers []rune // 运算符栈
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			// 找连续的数字
			j := i + 1
			for ; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
			}
			num, err := strconv.Atoi(s[i:j]) // 转换成整型
			if err != nil {
				break
			}
			i = j - 1

			// 运算符栈为空：直接插入
			if len(opers) == 0 {
				nums = append(nums, num)
			} else {
				// 提前处理乘除法
				oper := opers[len(opers)-1]
				top := nums[len(nums)-1]
				if oper == '*' {
					nums[len(nums)-1] = num * top
					opers = opers[:len(opers)-1]
				} else if oper == '/' {
					nums[len(nums)-1] = top / num
					opers = opers[:len(opers)-1]
				} else {
					nums = append(nums, num)
				}
			}
		} else if s[i] == '*' || s[i] == '/' { // 乘除法优先级高 直接插入
			opers = append(opers, rune(s[i]))
		} else if s[i] == '+' || s[i] == '-' { // 加减法优先级低 前面的运算符都需要计算
			if len(opers) == 0 {
				opers = append(opers, rune(s[i]))
			} else {
				oper := opers[len(opers)-1]
				top := nums[len(nums)-1]
				if oper == '*' {
					nums[len(nums)-2] *= top
				} else if oper == '/' {
					nums[len(nums)-2] /= top
				} else if oper == '+' {
					nums[len(nums)-2] += top
				} else {
					nums[len(nums)-2] -= top
				}
				opers[len(opers)-1] = rune(s[i])
				nums = nums[:len(nums)-1]
			}
		} else if s[i] == ' ' {
			continue
		}
		//fmt.Println("nums:", nums, " opers:", opers)
	}
	for len(opers) > 0 {
		oper := opers[len(opers)-1]
		top := nums[len(nums)-1]
		if oper == '*' {
			nums[len(nums)-2] *= top
		} else if oper == '/' {
			nums[len(nums)-2] /= top
		} else if oper == '+' {
			nums[len(nums)-2] += top
		} else if oper == '-' {
			nums[len(nums)-2] -= top
		}
		opers = opers[:len(opers)-1]
		nums = nums[:len(nums)-1]
	}
	return nums[len(nums)-1]
}

/**
 * 官方解答 更简洁
 * 把减法变成加法这种处理特别好，极大地简化了代码,省掉了操作符栈
 */
func calculate2(s string) (ans int) {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}
func main() {
	ss := []string{"3+2*2", " 3/2 ", " 3+5 / 2 ", "42", "33+20*22-5", "1-1+1", "14/3*2"}
	for _, s := range ss {
		fmt.Println(s, "=", calculate(s), calculate2(s))
	}
}
