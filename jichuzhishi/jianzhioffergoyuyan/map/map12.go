package main

import "fmt"

/*
数组中重复的数字
*/
func main() {
	nums := []int{ 1, 2, 2, 2, 3, 4, 5 }
	result := duplicate(nums)
	fmt.Println(result)
}

func duplicate(nums []int) bool{
	for i := 0; i < len(nums); i ++ {
		for j := i+1; j < len(nums); j ++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

