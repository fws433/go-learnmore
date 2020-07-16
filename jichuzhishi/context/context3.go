package main

import (
	"fmt"
	"sync"
	"time"
)

//方法二：如何优雅的退出子程序
//全局变量的方法
var wg1 sync.WaitGroup
var exit bool
func worker1(){
	for{
		fmt.Println("worker")
		time.Sleep(time.Second)
		if exit{
			break
		}
	}
	wg1.Done()
}
func main(){
	wg1.Add(1)
	go worker1()
	time.Sleep(time.Second*3)
	exit=true
	wg1.Wait()
	fmt.Println("over")
}
