package main

import (
	"fmt"
	"sort"
)

/*
最小k个数
*/

func main() {
	nums := []int{ 1, 2, 3, 4, 5, 6, 7 }
	k := 4
	result := getLeastNumbers(nums, k)
	fmt.Println(result)
}

func getLeastNumbers(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return nil
	}

	if k >= len(nums) {
		return nums
	}

	sort.Ints(nums)
	return nums[0 : k-1]
}
