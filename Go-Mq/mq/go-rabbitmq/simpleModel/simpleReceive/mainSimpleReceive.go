package main

import rabbitmq2 "mq/go-rabbitmq/simpleModel/rabbitmq"

func main() {
	rabbitmq :=rabbitmq2.NewRabbitMQSimple(""+"fws")
	rabbitmq.ConsumeSimple()
}
