package main

import "fmt"

func reverse(x int) int {
	v := make([]int, 0)
	tmp := x

	for tmp != 0 {
		v = append(v, tmp%10)
		tmp = tmp / 10
		//fmt.Println("reverse:", tmp)
	}

	var ret int = 0
	for i := 0; i < len(v); i++ {
		k := v[i]
		for j := 0; j < len(v)-i-1; j++ {
			k *= 10
		}
		if ret > 2147483647-k || ret < -2147483648-k {
			return 0
		}
		ret += k
		//fmt.Println("ret:", ret)
	}
	return ret
}
func reverse2(x int) int {
	tmp := x
	ret := 0
	for tmp != 0 {
		re := tmp % 10
		if re >= 0 {
			if ret > (2147483647-re)/10 {
				ret = 0
				break
			}
		} else {
			if ret < (-2147483648-re)/10 {
				ret = 0
				break
			}
		}
		//fmt.Println("reverse:", tmp)

		ret = 10*ret + re
		tmp = tmp / 10
	}
	return ret
}

func main() {
	var i int
	fmt.Scanf("%d", &i)
	fmt.Println("input:", i)

	fmt.Println("result:", reverse2(i))
}
