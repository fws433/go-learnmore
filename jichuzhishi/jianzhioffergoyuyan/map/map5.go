package main

import "fmt"

/*
数字在排序数组中出现的次数
*/
func main() {
	var nums = []int{1,2,2,2,3,4,5,6}
	k := 2
	count := getNumberOfk(nums, k)
	fmt.Println(count)
}

func getNumberOfk(nums []int, k int) int {
	length := len(nums)
	firstK := getFirstK(nums, k, 0, length-1)
	lastK := getLastK(nums, k, 0, length-1)
	if firstK <= lastK {
		return lastK - firstK + 1
	}
	return 0
}

func getLastK(nums []int, k int, start int, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2
	if nums[mid] > k {
		return getLastK(nums, k, start, mid-1)
	} else if nums[mid] < k {
		return getLastK(nums, k, mid+1, end)
	} else if nums[mid + 1] == k {
		return getLastK(nums, k, mid+1, end)
	} else {
		return mid
	}
}

func getFirstK(nums []int, k int, start int, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2
	if nums[mid] > k {
		return getFirstK(nums, k, start, mid-1)
	} else if nums[mid] < k {
		return getFirstK(nums, k, mid+1, end)
	} else if mid - 1 >= 0 && nums[mid-1] == k {
		return getFirstK(nums, k, start, mid-1)
	} else {
		return mid
	}
}
