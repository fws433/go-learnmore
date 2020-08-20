package main

import "fmt"

/*
数组中出现次数超过一半度数字
*/
func main(){
	nums := []int{ 1, 2, 2, 2, 2, 3, 3, 4 }
	result := moreThanHalfNum(nums)
	fmt.Println(result)
}

func moreThanHalfNum(nums []int) int {
	if nums == nil {
		return 0
	}

	count := 0
	length := len(nums)
	//sort.Ints(nums)

	for i := 0; i < length; i ++ {
		for j :=0; j< length; j ++ {
			if nums[i] == nums[j] {
				count ++
			}
		}

		if count >= length / 2 {
			return nums[i]
		}
	}
	return 0
}
