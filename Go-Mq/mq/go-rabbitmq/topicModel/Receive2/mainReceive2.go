package main

import "go-rabbitmq/go-rabbitmq/topicModel/rabbitMQ"

func main() {
	fwsTwo :=rabbitMQ.NewRabbitMQTopic("fwstopic", "fws.*.two")
	fwsTwo.ReceiveTopic()
}
