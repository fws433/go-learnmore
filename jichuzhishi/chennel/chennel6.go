package main

import (
	"fmt"
	"sync"
)
//主线程为了等待goroutine都运行完毕，不得不在程序的末尾使用time.Sleep() 来睡眠一段时间，等待其他线程充分运行。对于简单的代码，100个for循环可以在1秒之内运行完毕，time.Sleep() 也可以达到想要的效果。
//
//但是对于实际生活的大多数场景来说，1秒是不够的，并且大部分时候我们都无法预知for循环内代码运行时间的长短。这时候就不能使用time.Sleep() 来完成等待操作了。
func main(){
	wg:=sync.WaitGroup{}
	wg.Add(100)
	 for i:=0;i<100;i++{
	 	go func(i int){
			fmt.Println(i)
			wg.Done()
	 	}(i)
	}
	wg.Wait()
}
