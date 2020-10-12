package main

import "fmt"

/**
 * LC 删除排序数组中的重复项
 * https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2gy9m/
 */

func removeDuplicates(nums []int) int {
	var size = len(nums)
	if size <= 1 {
		return len(nums)
	}
	var p = 0
	for q := 1; q < size; q++ {
		if nums[p] != nums[q] {
			if q-p > 1 {
				nums[p+1] = nums[q]
			}
			p++
		}
	}
	return p + 1
}

func main() {
	// var a1 = []int{1, 1, 2}
	var a1 = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	var s = removeDuplicates(a1)

	fmt.Println("Source Array", a1)
	fmt.Println("New Array Len", s)
	fmt.Println("New Array", a1[:s])
}
