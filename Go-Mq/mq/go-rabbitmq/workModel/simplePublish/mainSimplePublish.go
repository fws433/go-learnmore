package main

import (
	"fmt"
	"strconv"
	"time"

	"go-rabbitmq/go-rabbitmq/workModel/rabbitmq"
)

func main() {
 	rabbitmq :=rabbitmq.NewRabbitMQSimple(""+"fws")
 	for i := 0;i <= 100;i++ {
 		rabbitmq.PublishSimple("hello fws" + strconv.Itoa(i))
 		time.Sleep(1 * time.Second)
 		fmt.Println(i)
	}
 }
