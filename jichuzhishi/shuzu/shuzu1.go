package main

import "fmt"

func main(){
	//定义一个长度为10的数组
	var n [10]int
	var i,j int
	//为数组初始化
	for i=0;i<10;i++{
		n[i]=i+100
	}
	//输出每个数组中的元素
	for j=0;j<10;j++{
		fmt.Println(n[j])
	}
}
