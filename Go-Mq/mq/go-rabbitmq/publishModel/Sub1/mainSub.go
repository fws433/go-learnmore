package main

import (
	"go-rabbitmq/go-rabbitmq/publishModel/rabbitmq"
)

func main() {
	rabbitmq :=rabbitmq.NewRabbitMQPubSub("logs")
	rabbitmq.ReceiveSub()
}
