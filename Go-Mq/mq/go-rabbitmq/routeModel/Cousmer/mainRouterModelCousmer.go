package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err !=nil {
		log.Fatalf("%s:%s", msg,err)
	}
}

func main() {
	conn, err :=amqp.Dial("amqp://fws:guest@127.0.0.1:5672/fws")
	failOnError(err, "Failed to connect to rabbitmq")
	defer conn.Close()

	ch, err :=conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	//尝试创建交换机
	err = ch.ExchangeDeclare(
		"logs",
		"direct",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Failed to declare a queue")

	//2.试探性创建队列
	q, err :=ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
		)

	if len(os.Args) <2 {
		log.Printf("Usage: %s [info] [warning] [error]",os.Args[0])
		os.Exit(0)
	}
	for _,s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, "log", s)
		err =ch.QueueBind(
			q.Name,
			s,
			"log",
			false,
			nil,
			)
	}

	msgs, err :=ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "Failed to register a consumer")

	forever :=make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("[x] %s",d.Body)
		}
	}()

	log.Printf("[*] waiting for logs. To exit press ctrl+c")
	<-forever
}






