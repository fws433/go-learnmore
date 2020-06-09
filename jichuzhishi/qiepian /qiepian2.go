package main

import "fmt"

func main(){
	a:=[5]int{1,2,3,4,6}
	s:=a[1:3]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n",s,len(s),cap(s))
	fmt.Println(s[1])
	s2:=s[3:4]
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n",s2,len(s2),cap(s2))
}