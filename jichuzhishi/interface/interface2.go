package main

import "fmt"

//一个类型实现多个接口
type Sayer interface{
	say()
}
//Mover接口
type Mover interface {
	move()
}
//dog既可以实现Sayer接口，也可以实现Mover接口
type dog struct{
	name string
}
//实现Sayer接口
func (d dog)say(){
	fmt.Printf("%s会叫汪汪汪\n",d.name)
}
//实现mover接口
func (d dog)move(){
	fmt.Printf("%s会动\n",d.name)
}
func main(){
	var x Sayer
	var y Mover
	var a =dog{
		name:"摩托"}
	x=a
	y=a
	x.say()
	y.move()
}



