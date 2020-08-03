package main

import (
	"fmt"
	"strconv"
	"time"

	"go-rabbitmq/go-rabbitmq/publishModel/rabbitmq"
)

func main() {
	rabbitmq1 :=rabbitmq.NewRabbitMQPubSub(""+"newProduct")
	for i :=0; i < 100; i++ {
		rabbitmq1.PublishPub("订阅模式生产第" + strconv.Itoa(i) + "条" + " 数据")
		fmt.Println("订阅模式生产者" + strconv.Itoa(i) + "条" +"数据")
		time.Sleep(1 * time.Second)
	}
}
