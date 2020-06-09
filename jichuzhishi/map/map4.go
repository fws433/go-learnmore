package main

import "fmt"

//元素为map类型的切片
func main(){
	var slicemap=make([]map[string]string,3)
	for index,value:=range slicemap{
		fmt.Printf("index:%d value:%v\n",index,value)
	}
	fmt.Println("after init")
	//对切片中的map元素进行初始化
	slicemap[0]=make(map[string]string,10)
	slicemap[0]["name"]="小王子"
	slicemap[0]["password"]="12345"
	slicemap[0]["address"]="沙河"
	for index,value:=range slicemap{
		fmt.Printf("index:%d value:%v\n",index,value)
	}
}
