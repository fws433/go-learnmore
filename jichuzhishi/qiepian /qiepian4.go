package main

import "fmt"

//切片的复制
func main(){
	a:=[]int{1,2,3,4,5}
	b:=a//切片ab指向同一内存地址，修改a的同时b的值也会发生变化
	fmt.Println(a)
	fmt.Println(b)
	b[0]=1000
	fmt.Println(a)
	fmt.Println(b)
	//使用copy将切片a中的元素复制到切片c
	c:=make([]int,5,5)
	copy(c,a)
	fmt.Println(a)
	fmt.Println(c)
	a[0]=1000
	c[0]=2000
	fmt.Println(a)
	fmt.Println(c)
	//删除索引
	d:=[]int{30,31,32,33,35}
	//删除索引为2 的元素
	d=append(d[:2],d[3:]...)
	fmt.Println(d)

}
