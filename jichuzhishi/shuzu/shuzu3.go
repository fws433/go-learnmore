package main

import "fmt"

//让编译器根据初始值的个数自行推断数组的长度
func main(){
	var testarray [3]int
	var numarray=[...]int{1,2}
	var cityarray=[...]string{"北京","上海","深圳"}
	fmt.Println(testarray)
	fmt.Println(numarray)
	fmt.Printf("type of numArray:%T\n",numarray)//数组的类型
	fmt.Println(cityarray)
	fmt.Printf("type of cityarray:%T\n",cityarray)//数组类型
}