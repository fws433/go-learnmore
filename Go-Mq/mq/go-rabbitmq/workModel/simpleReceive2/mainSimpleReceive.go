package main

import "go-rabbitmq/go-rabbitmq/workModel/rabbitmq"

func main() {
	rabbitmq :=rabbitmq.NewRabbitMQSimple(""+"fws")
	rabbitmq.ConsumeSimple()
}
