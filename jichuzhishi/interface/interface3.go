package main

import "fmt"

//多个类型实现同一接口

//Mover接口
type Mover1 interface {
	move()
}
type car struct{
	brand string
}
type dog1 struct{
	name string
}
//dog类型实现mover接口
func(d dog1)move(){
	fmt.Printf("%s会跑\n",d.name)
}
func (c car)move(){
	fmt.Printf("%s速度70\n",c.brand)
}
func main(){
	var x Mover1
	var a=dog1{name:"大大"}
	var b1=car{brand:"保时捷"}
	x=a
	x.move()
	x=b1
	x.move()


}

