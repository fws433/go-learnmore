package main

import "fmt"

func main(){
	//append(）添加元素及扩容
	var numslice []int
	for i:=0;i<10;i++{
		numslice=append(numslice,i)
		fmt.Printf("%v len:%d cap:%d ptr:%p\n",numslice,len(numslice),cap(numslice),numslice)
	}
	var cityslice []string
	//追加一个元素
	cityslice=append(cityslice,"北京")
	//追加多个切片
	cityslice=append(cityslice,"上海","广州","深圳")
	//追加切片
	a:=[]string{"成都","重庆"}
	cityslice=append(cityslice,a...)
	fmt.Println(cityslice)

}
