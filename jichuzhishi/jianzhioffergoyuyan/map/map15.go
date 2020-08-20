package main

import "fmt"

/*
二维数组的查找
*/
func main() {
   nums := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}, {4, 5, 6}}
   target := 7
   bl := searchMatrix(nums, target)
   fmt.Println(bl)
}

func searchMatrix(nums [][]int, target int) bool {
	if nums == nil || len(nums) < 1 {
		return false
	}
	row := 0
	col := len(nums[0]) - 1
	fmt.Println(len(nums))
	fmt.Println(len(nums[0]))
	for row <= len(nums) - 1 && col >= 0{
		 if nums[row][col] > target {
		 	col --
		 } else if nums[row][col] < target {
		 	row ++
		 } else {
		 	return true
		 }
	}
	return false
}

