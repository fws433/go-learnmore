package main

import "fmt"

/*
输入一个递增排序的数组和一个数s,在数组中查找两个数，使得他们的和正和是s,
如果有多对数字的和等于s,输出两个数的乘极是最小的
*/
func main() {
	var nums =[...]int {1,2,3,4,5,6}
	var num = 6
	result := findNumbersWithSum(nums, num)
	fmt.Println(result)
}

func findNumbersWithSum(a [6]int, num int)[] int {
	result := []int{}
	length :=len(a)
    value :=0

	if length == 0 {
		return  result
	}

	for i := 0 ; i < length- 1 ; i ++ {
		for j := i + 1 ; j < length -1 ;j ++ {
			if a [i] + a [j] == num {
				if value != 0 {
					if value > (a [i] * a [j]) {
						value =a [i] * a [j]
						result =append (result, a[i] ,a[j])
					}
				} else {
					value = a [i] * a [j]
					result = append (result, a[i], a[j])
				}
			}
		}
	}
	return result

}
