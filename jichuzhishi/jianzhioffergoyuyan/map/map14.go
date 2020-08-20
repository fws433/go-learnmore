package main

import (
	"fmt"
	"sort"
)
/*数据流中的中位数*/
func main() {
	nums := []int{ 1, 2, 3, 4, 5, 6 }
	result := getMedian(nums)
	fmt.Println(result)
}

func getMedian(nums []int) float64 {
	sort.Ints(nums)
	length := len(nums)
	if length == 0{
		return -1
	}

	sort.Ints(nums)
	if length % 2 ==0 {
		return float64(nums[length/2-1] + nums[length/2])/2
	} else {
		return float64(nums[length/2])
	}
}