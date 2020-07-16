package main

import (
	"fmt"
	"sync"
	"time"
)

//方法1：如何优雅的结束子goroutine
var wg sync.WaitGroup
func worker(){
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	wg.Done()
}
func main(){
	wg.Add(1)
	go worker()
	wg.Wait()
	fmt.Println("over")
}