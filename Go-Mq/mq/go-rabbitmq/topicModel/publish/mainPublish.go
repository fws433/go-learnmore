package main

import (
	"fmt"
	"strconv"
	"time"

	"go-rabbitmq/go-rabbitmq/topicModel/rabbitMQ"
)

func main() {
	fwsOne :=rabbitMQ.NewRabbitMQTopic("fwstopic", "fws.topic.one")
	fwsTwo :=rabbitMQ.NewRabbitMQTopic("fwstopic", "fws.topic.two")
	for i :=0; i <=100; i++ {
		fwsOne.PublishTopic("hello fws topic one!" + strconv.Itoa(i))
		fwsTwo.PublishTopic("hello fws topic two!" + strconv.Itoa(i))
		time.Sleep(1*time.Second)
		fmt.Println(i)
	}
}
