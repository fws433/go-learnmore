package main

import rabbitmq2 "mq/go-rabbitmq/publishModel/rabbitmq"

func main() {
	rabbitmq :=rabbitmq2.NewRabbitMQPubSub(""+"newProduct")
	rabbitmq.ReceiveSub()
}
