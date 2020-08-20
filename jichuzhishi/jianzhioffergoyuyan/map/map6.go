package main

import "fmt"

/*
数组中只出现一次的数字
*/
func main() {
	nums :=[]int{1, 2, 2, 3, 3, 4, 4}
	result := singlerNumber(nums)
	fmt.Println(result)
}

func singlerNumber(nums []int) int {
	result := 0
	for _, num :=range nums {
		result = result^num
	}
	return result
}

