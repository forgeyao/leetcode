package main

import "testing"

func TestAddDigits(t *testing.T) {
	nums := []int{38, 0}
	results := []int{2, 0}
	for i := 0; i < len(nums) && i < len(results); i++ {
		ret := AddDigits(nums[i])
		if ret != results[i] {
			t.Fatal("num:", nums[i], "ret:", ret, "result:", results[i])
		}
	}
}
