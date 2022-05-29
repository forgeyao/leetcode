package main

import (
	"fmt"
	"strconv"
)

/**
 * 倒序遍历两个字符串，依次相加相对的字符，结果存在 slice
 * 最后反转 slice 即为结果
 * 时间 O(m) m为num1/num2中较长的字符串长度
 * 空间 O(1)
 */
func addStrings(num1 string, num2 string) string {
	var ret []byte
	var flag byte
	i, j := len(num1)-1, len(num2)-1
	for ; i >= 0 || j >= 0; i, j = i-1, j-1 {
		tmp := flag
		if i >= 0 {
			tmp += num1[i] - '0'
		}
		if j >= 0 {
			tmp += num2[j] - '0'
		}
		if tmp > 9 {
			flag = 1
			ret = append(ret, tmp-10+'0')
		} else {
			flag = 0
			ret = append(ret, tmp+'0')
		}
	}
	if flag == 1 {
		ret = append(ret, flag+'0')
	}
	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}
	return string(ret)
}

/**
 * 参考的官方答案。更简洁
 */
func addStrings2(num1 string, num2 string) string {
	var ret string
	var flag byte
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || flag != 0; i, j = i-1, j-1 {
		tmp := flag
		if i >= 0 {
			tmp += num1[i] - '0'
		}
		if j >= 0 {
			tmp += num2[j] - '0'
		}
		flag = tmp / 10
		ret = strconv.Itoa(int(tmp%10)) + ret
	}
	return string(ret)
}

func main() {
	num1 := []string{"11", "456", "0", "1"}
	num2 := []string{"123", "77", "0", "9"}
	for i := 0; i < len(num1) && i < len(num2); i++ {
		fmt.Println(num1[i], num2[i], " ret:", addStrings(num1[i], num2[i]), addStrings2(num1[i], num2[i]))
	}
}
