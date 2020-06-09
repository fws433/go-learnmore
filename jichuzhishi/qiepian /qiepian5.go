package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		//a = append(a, fmt.Sprintf("%v", i))
		a=append(a,strconv.Itoa(i))
	}
	fmt.Println(a)
}
