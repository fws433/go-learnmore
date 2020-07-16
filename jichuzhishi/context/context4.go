package main

import (
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup
func worker2(exitChan chan struct{}){
	LOOP:
		for{
			fmt.Println("worker")
			time.Sleep(time.Second)
			select {
				case<-exitChan://等待接收上级通知
				    break LOOP
				default:
			}
		}
		wg2.Done()
}
func main(){
	var exitChan=make(chan struct{})
	wg2.Add(1)
	go worker2(exitChan)
	time.Sleep(time.Second*3)
	exitChan<- struct{}{} //给子goroutine发送退出信号
	wg2.Wait()
	fmt.Println("over")
}
