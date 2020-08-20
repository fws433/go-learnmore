package main

import "fmt"
/*
1。给定1-100的整数数组，找到其中缺少（重复）的值

*/
func main() {
	var nums = [...]int{1,8,8}
	repeatOrLessNum(nums)
}

func repeatOrLessNum(nums [3]int) {
	var base [101]int
	for i := 0; i < len(nums); i++ {
		base[nums[i]] = nums [i]
	}

	for j := 0; j < len(base); j ++ {
		if base[j] == j {
			continue
		} else {
			fmt.Println("缺少的数字" ,+j )
		}
	}
}
