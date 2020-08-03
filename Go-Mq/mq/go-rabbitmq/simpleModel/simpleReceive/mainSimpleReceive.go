package main

import (
	"go-rabbitmq/go-rabbitmq/simpleModel/rabbitmq"
)

func main() {
	rabbitmq :=rabbitmq.NewRabbitMQSimple(""+"fws")
	rabbitmq.ConsumeSimple()
}
