package main

import "fmt"

/*
连续子数组的最大和
*/

func main(){
	nums := []int{6, -3, -2, 7, -15, 1, 2, 2}
	result := maxSubArray(nums)
	fmt.Println(result)
}

func maxSubArray(nums []int) int {
	sum := 0
	res := nums[0]
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num
		}

		if res < sum {
			res = sum
		}
	}
	return  res
}
