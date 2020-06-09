package main

import "fmt"

// Sayer 接口


// 接口嵌套
type animal interface {
	say()
	move()
}
type cat3 struct {
	name string
}
/*type cat4 struct{
	cat3
}*/

func (c cat3) say() {
	fmt.Println("喵喵喵")
}


func (c cat3) move() {
	fmt.Println("猫会动")
}

func main() {
	var x animal
	x = cat3{name: "花花"}
	x.move()
	x.say()
}