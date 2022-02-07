//https://leetcode-cn.com/problems/longest-happy-string/
package main

import (
	"fmt"
	"sort"
)

/**
 * "贪心"法
 * 按顺序，先尽可能多放各个元素，每次追加在 slice 最后面，直到有两个元素都变为0(再继续追加剩余的元素会出现3个连续)
 * 再遍历 slice，找可继续插入的位置，插入元素
 * 最后，看 slice 是否还能补充元素
 * 时间 O(a+b+c)
 * 空间 O(a+b+c), tmp 缓存临时 slice
 */
func longestDiverseString(a int, b int, c int) string {
	ret := []byte{}
	for a > 0 && b > 0 || a > 0 && c > 0 || b > 0 && c > 0 {
		for i := 0; i < 2 && a > 0; i++ {
			ret = append(ret, 'a')
			a--
		}
		for i := 0; i < 2 && b > 0; i++ {
			ret = append(ret, 'b')
			b--
		}
		for i := 0; i < 2 && c > 0; i++ {
			ret = append(ret, 'c')
			c--
		}
	}
	for i := 0; i < len(ret); i++ {
		if a > 0 {
			if i == 0 && ret[i] != 'a' || i > 0 && ret[i] != 'a' && ret[i-1] != 'a' {
				tmp := []byte{}
				for j := 0; j < 2 && a > 0; j++ {
					tmp = append(tmp, 'a')
					a--
				}
				j := len(tmp)
				//fmt.Println(i, "tmp:", string(tmp), "ret[i:]:", string(ret[i:]))
				tmp = append(tmp, ret[i:]...)
				ret = append(ret[:i], tmp...)
				i += j
			}
			//fmt.Println(i, string(ret))
		}
		if b > 0 {
			if i == 0 && ret[i] != 'b' || i > 0 && ret[i] != 'b' && ret[i-1] != 'b' {
				tmp := []byte{}
				for j := 0; j < 2 && b > 0; j++ {
					tmp = append(tmp, 'b')
					b--
				}
				j := len(tmp)
				tmp = append(tmp, ret[i:]...)
				ret = append(ret[:i], tmp...)
				i += j
			}
		}
		if c > 0 {
			if i == 0 && ret[i] != 'c' || i > 0 && ret[i] != 'c' && ret[i-1] != 'c' {
				tmp := []byte{}
				for i := 0; i < 2 && c > 0; i++ {
					tmp = append(tmp, 'c')
					c--
				}
				j := len(tmp)
				//fmt.Println(i, "tmp:", string(tmp), "ret[i:]:", string(ret[i:]))

				tmp = append(tmp, ret[i:]...)
				ret = append(ret[:i], tmp...)
				i += j
			}
			//fmt.Println(i, string(ret))
		}
		if a == 0 && b == 0 && c == 0 {
			break
		}
	}
	if a > 0 && (len(ret) == 0 || len(ret) > 0 && ret[len(ret)-1] != 'a') {
		for j := 0; j < 2 && a > 0; j++ {
			ret = append(ret, 'a')
			a--
		}
	}
	if b > 0 && (len(ret) == 0 || len(ret) > 0 && ret[len(ret)-1] != 'b') {
		for j := 0; j < 2 && b > 0; j++ {
			ret = append(ret, 'b')
			b--
		}
	}
	if c > 0 && (len(ret) == 0 || len(ret) > 0 && ret[len(ret)-1] != 'c') {
		for j := 0; j < 2 && c > 0; j++ {
			ret = append(ret, 'c')
			c--
		}
	}
	return string(ret)
}

func f(a *int, b byte, ret *[]byte) {
	for i := 0; i < 2 && *a > 0; i++ {
		*ret = append(*ret, b)
		*a--
	}
}

/**
 * 思路跟方法一一样，只有优化重复的代码
 */
func longestDiverseString2(a int, b int, c int) string {
	ret := []byte{}
	for a > 0 && b > 0 || a > 0 && c > 0 || b > 0 && c > 0 {
		f(&a, 'a', &ret)
		f(&b, 'b', &ret)
		f(&c, 'c', &ret)
	}

	for i := 0; i < len(ret); i++ {
		tmp := []byte{}
		if i == 0 && ret[i] != 'a' || i > 0 && ret[i] != 'a' && ret[i-1] != 'a' {
			f(&a, 'a', &tmp)
		}
		if i == 0 && ret[i] != 'b' || i > 0 && ret[i] != 'b' && ret[i-1] != 'b' {
			f(&b, 'b', &tmp)
		}
		if i == 0 && ret[i] != 'c' || i > 0 && ret[i] != 'c' && ret[i-1] != 'c' {
			f(&c, 'c', &tmp)
		}
		if len(tmp) == 0 {
			continue
		}
		j := len(tmp)
		tmp = append(tmp, ret[i:]...)
		ret = append(ret[:i], tmp...)
		i += j

		if a == 0 && b == 0 && c == 0 {
			break
		}
	}
	if len(ret) == 0 || len(ret) > 0 && ret[len(ret)-1] != 'a' {
		f(&a, 'a', &ret)
	}
	if len(ret) == 0 || len(ret) > 0 && ret[len(ret)-1] != 'b' {
		f(&b, 'b', &ret)
	}
	if len(ret) == 0 || len(ret) > 0 && ret[len(ret)-1] != 'c' {
		f(&c, 'c', &ret)
	}
	return string(ret)
}

/**
 * 官方解答
 * 很巧妙的用了排序，代码更简洁（我有想到排序，但没想到构造 struct 来实现排序）
 * 时间 O((a+b+c)* C*logC) C 表示字母种类，C*logC 表示排序
 * 空间 O(C)
 */
func longestDiverseString3(a int, b int, c int) string {
	ans := []byte{}
	cnt := []struct {
		c  int
		ch byte
	}{{a, 'a'}, {b, 'b'}, {c, 'c'}}
	for {
		sort.Slice(cnt, func(i, j int) bool { return cnt[i].c > cnt[j].c })
		hasNext := false
		for i, p := range cnt {
			if p.c <= 0 {
				break
			}
			m := len(ans)
			if m >= 2 && ans[m-2] == p.ch && ans[m-1] == p.ch {
				continue
			}
			hasNext = true
			ans = append(ans, p.ch)
			cnt[i].c--
			break
		}
		if !hasNext {
			return string(ans)
		}
	}
}

func main() {
	as := []int{1, 2, 7, 0}
	bs := []int{1, 2, 1, 0}
	cs := []int{7, 1, 0, 7}
	for i := 0; i < len(as) && i < len(bs) && i < len(cs); i++ {
		fmt.Println(as[i], bs[i], cs[i], "ret:", longestDiverseString(as[i], bs[i], cs[i]), longestDiverseString2(as[i], bs[i], cs[i]), longestDiverseString3(as[i], bs[i], cs[i]))
	}
}
