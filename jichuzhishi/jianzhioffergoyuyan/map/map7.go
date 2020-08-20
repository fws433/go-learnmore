package main

import "fmt"

/*
旋转数组的最小数字  (有点难度)
*/
func main() {
	nums := []int{ 3, 4, 5, 1, 2 }
	result := minNumberInRotateArray(nums)
	fmt.Println(result)
}

func minNumberInRotateArray(nums []int) int {
	high := len(nums)-1
	low := 0

	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[high]{
			low = mid + 1
		}
		if nums[mid] < nums[high] {
			high = mid
		}
		if nums[mid] == nums[high] {
			high -= 1
		}
	}
	return nums[low]
}