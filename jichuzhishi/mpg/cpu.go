package main

import (
	"fmt"
	"runtime"
)

func main(){
	//获取当前系统逻辑cpu的数量
	num:=runtime.NumCPU()
	//可以设置使用的cpu逻辑核心个数，1.8之前
	runtime.GOMAXPROCS(num)
	fmt.Println("num=",num)

}
