package main

import (
	"fmt"
	"time"
)

func main(){
	t,err:=time.Parse("2006-01-02 15:04:05","2018-02-04 12:00:00")
	if err!=nil{
		fmt.Println(err)
		return
	}
	go func() {}()
	fmt.Println(t)
	fmt.Println(time.Now())
}
