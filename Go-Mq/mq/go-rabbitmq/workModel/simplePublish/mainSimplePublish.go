package main

import (
	"fmt"
	"strconv"
	"time"

	rabbitmq2 "mq/go-rabbitmq/simpleModel/rabbitmq"
)

func main() {
 	rabbitmq :=rabbitmq2.NewRabbitMQSimple(""+"fws")
 	for i := 0;i <= 100;i++ {
 		rabbitmq.PublishSimple("hello fws" + strconv.Itoa(i))
 		time.Sleep(1 * time.Second)
 		fmt.Println(i)
	}
 }
