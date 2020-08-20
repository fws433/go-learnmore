package main
func main() {
	nums := [][]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {6, 7, 8, 9}, {10, 11, 12, 13}, {14, 15, 16, 17}}
	printMatrix(nums)

}
func printMatrix(nums [][]int) []int {
	result :=[]int{}
	//从上到下
	row := 0
	column := len(nums)
	if column == 0 {
		return result
	}
	//从左到右
	left := 0
	if len(nums[0]) == 0 {
		return  result
	}
	right := len(nums[0])-1
	for left <= right && row <= column-1 {
		for i := left; i <= right; i ++ {
			result = append(result, nums[row][i])
		}
		for i := row + 1; i < column; i ++ {
			result = append(result, nums[i][right])
		}
		for i := right - 1; i >= left; i-- {
			result = append(result, nums[i][left])
		}
		for i := column-2; i > row; i -- {
			result = append(result, nums[i][left])
		}
		left ++
		right --
		row ++
		column --
	}
	return  result
}