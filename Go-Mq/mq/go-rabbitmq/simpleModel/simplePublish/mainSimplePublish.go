package main

import (
	"fmt"

	"go-rabbitmq/go-rabbitmq/simpleModel/rabbitmq"
)

func main() {
 	rabbitmq :=rabbitmq.NewRabbitMQSimple(""+"fws")
 	rabbitmq.PublishSimple("hello world")
 	fmt.Println("发送成功！")
 }
