package main

import "fmt"

// 10021 1000030001
func isPalindromeOld(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	tmp := x
	ten := 1
	for tmp /= 10; tmp != 0; {
		ten *= 10
		tmp /= 10
	}

	tmp = x
	for tmp > 9 {
		if tmp/ten != tmp%10 {
			return false
		}
		tmp = (tmp%ten - tmp%10) / 10
		ten /= 100
		if ten > 10 && (tmp > 0 && tmp < 10) {
			return false
		}
	}
	return true
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	tmp := x
	ten := 1
	length := 1
	for tmp /= 10; tmp != 0; length++ {
		ten *= 10
		tmp /= 10
	}
	fmt.Println("ten:", ten, "len:", length)

	tmp = x
	high := length - 1
	low := 0
	for high > low {
		fmt.Println("high:", high, "low:", low, "tmp:", tmp, "ten:", ten)
		if tmp/ten != tmp%10 {
			return false
		}
		tmp = (tmp%ten - tmp%10) / 10
		ten /= 100
		high--
		low++
	}
	return true
}
