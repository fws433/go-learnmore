package main

import "fmt"

/*
和为s的连续正数序列
*/
func main() {
	sum := 100
	result := findContinueOusSequence(sum)
	fmt.Println(result)
}

func findContinueOusSequence(sum int) [][]int{
	result := [][]int{}
	low := 1
	high := 2
	for low < high {
		temp := (low + high)*(high-low+1)/2
		if temp == sum {
			cur := []int{}
			for i := low; i <= high; i ++ {
				cur = append(cur, i)
			}
			result = append(result, cur)
			low ++
		}

		if temp < sum {
			high ++
		}

		if temp > sum {
			low ++
		}
	}

	return result
}