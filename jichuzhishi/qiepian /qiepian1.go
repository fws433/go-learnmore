package main

import "fmt"

func main(){
	//创建切片
	num:=[]int{1,2,3,4,5,6}
	printslice(num)
	//打印原始切片
	fmt.Println("num=",num)
}

func printslice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}
