package main

import "fmt"

func main(){
	a:="how do you do"
	word:=make([]string,4)
	fmt.Println(word)
	counter:=0
	how:=0
	do:=0
	you:=0
	for i:=range a{
		value:=fmt.Sprintf("%c",a[i])
		if value!=" "{
			word[counter]+=value
		}else{
			counter++
		}
	}
	for _,v:=range word{
		switch  v {
		case "how":
			how ++
		case "do":
			do++
		case "you":
			you++
		}
	}
	fmt.Println(word)
	fmt.Printf("how=%d\t",how)
	fmt.Printf("do=%d\t",do)
	fmt.Printf("you=%d\t",you)




}
