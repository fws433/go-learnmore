package main

import "go-rabbitmq/go-rabbitmq/topicModel/rabbitMQ"

func main() {
	fwsOne :=rabbitMQ.NewRabbitMQTopic("fwstopic", "#")
	fwsOne.ReceiveTopic()
}
