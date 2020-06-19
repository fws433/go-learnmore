package main

import (
	"fmt"
	"sync"
	"time"
)

//sync.once.do(f func())能保证once只执行一次，无论你是否更换once.do(xx)这里的方法，这个sync.once只会执行一次
var once sync.Once
func main(){
	for i,v:=range make([]string,10){
	once.Do(onces)
	fmt.Println("count:",v,"---",i)
}
for i:=0;i<10;i++{
	go func(){
		once.Do(onced)
		fmt.Println("213")
	}()
  }
  time.Sleep(4000)
}
func onced(){
	fmt.Println("onced")
}
func onces(){
	fmt.Println("onces")
}