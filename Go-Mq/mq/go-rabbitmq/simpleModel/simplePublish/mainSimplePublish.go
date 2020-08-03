package main

import (
	"fmt"

	rabbitmq2 "mq/go-rabbitmq/simpleModel/rabbitmq"
)

func main() {
 	rabbitmq :=rabbitmq2.NewRabbitMQSimple(""+"fws")
 	rabbitmq.PublishSimple("hello world")
 	fmt.Println("发送成功！")
 }
