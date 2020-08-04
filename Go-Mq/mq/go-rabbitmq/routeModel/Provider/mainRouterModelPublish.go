package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func failOnError (err error ,msg string) {
	if err !=nil {
		log.Fatalf("%s: %s",msg,err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://fws:guest@127.0.0.1:5672/fws")
	failOnError(err, "Failed to connect to RabbitMQ")
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
	failOnError(err, "Failed to declare an exchange")

	body :=bodyForm(os.Args)
	err =ch.Publish(
		"logs",
		severityFrom(os.Args),
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		  })

	failOnError(err, "Failed to publish a message")

	log.Printf("[x] Sent %s",body)
}

func bodyForm(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], "")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "info"
	} else {
		s=os.Args[1]
	}
	return s
}









