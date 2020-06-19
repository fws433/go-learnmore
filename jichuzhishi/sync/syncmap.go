package main

import (
	"fmt"
	"strconv"
	"sync"
)

//go语言中内置的map不是并发安全的
//var m=make(map[string]int)
var m=sync.Map{}
func main(){
	wg:=sync.WaitGroup{}
	for i:=0;i<20;i++{
		wg.Add(1)
		go func(n int){
			key:=strconv.Itoa(n)
			//set(key,n)
			m.Store(key,n)
			value,_:=m.Load(key)
			fmt.Printf("K:=%v,V:=%v\n",key,value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
/*func get(key string)int{
	return m[key]
}
func set(key string,value int){
	m[key]=value
}*/
