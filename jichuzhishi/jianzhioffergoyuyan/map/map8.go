package main

import "fmt"

/*
逆序对（此方法不严谨，算法复杂度高）
*/
func main() {
	nums := []int{ 1, 2, 3, 4, 5, 6, 7, 0 }
	result := inversePairs(nums)
	fmt.Println(result)
}

func inversePairs(nums []int) int  {
	count := 0
	for i := 0; i < len(nums); i ++ {
		for j := i + 1; j < len(nums); j ++ {
			if nums[i] > nums[j] {
				count ++
			}
		}
	}

	return count % 10000007
}
