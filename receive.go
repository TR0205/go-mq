package main

import (
  "log"

  amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}

func main() {
	// MQと接続
	conn, err := amqp.Dial("amqp://user:password@mq:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// チャンネル経由でMQの機能をAPIとして使用可能
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// Publisherより先に開始することもあるため、Consumerにもキューを定義する
	q, err := ch.QueueDeclare(
	"hello", // name
	false,   // durable
	false,   // delete when unused
	false,   // exclusive
	false,   // no-wait
	nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	
	var forever chan struct{}
	
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}