package main

import (
	"fmt"
	"strconv"
)

/*
将数组排成最小的数，例如输入数组{3，32，321}，则打印出这三个数字能排成的最小数字为321323。
*/
func main(){
	nums := []int{ 3, 32, 321 }
	result := printMinNumber(nums)
	fmt.Println(result)
}

func printMinNumber(nums []int) string {
	for i := 0; i < len(nums); i ++ {
		for j := i+1; j < len(nums); j ++ {
			//int转string
			m := strconv.Itoa( nums[i] ) + strconv.Itoa( nums[j] )
			n := strconv.Itoa( nums[j] ) + strconv.Itoa( nums[i] )
			//string转int
			mm, _ := strconv.Atoi( m )
			nn, _ := strconv.Atoi( n )
			if mm > nn {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	str := ""
	for _, v := range nums {
		str += strconv.Itoa(v)
	}
	return str

}
