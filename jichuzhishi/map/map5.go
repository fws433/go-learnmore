package main

import "fmt"

//值为切片类型的map
func main(){
	var slicemap=make(map[string][]string,3)
	fmt.Println(slicemap)
	fmt.Println("after init")
	key:="🇨🇳"
	value,ok:=slicemap[key]
	if !ok{
		value=make([]string,0,2)
	}
	value=append(value,"北京","上海")
	slicemap[key]=value
	fmt.Println(slicemap)

}
