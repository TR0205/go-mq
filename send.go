package main

import (
  "context"
  "log"
  "time"

  amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
  // RabbitMQコンテナへ接続
  conn, err := amqp.Dial("amqp://user:password@mq:5672/")
  failOnError(err, "Failed to connect to RabbitMQ")
  defer conn.Close()

  // チャンネルの作成 チャンネル経由でMQの機能をAPIとして使用可能
  ch, err := conn.Channel()
  failOnError(err, "Failed to open a channel")
  defer ch.Close()

  // キューの定義
  q, err := ch.QueueDeclare(
    "hello", // name
    false,   // durable
    false,   // delete when unused
    false,   // exclusive
    false,   // no-wait
    nil,     // arguments
  )
  failOnError(err, "Failed to declare a queue")
  // タイムアウト設定
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()

  // メッセージの送信
  body := "Hello World!"
  err = ch.PublishWithContext(ctx,
    "",     // exchange
    q.Name, // routing key
    false,  // mandatory
    false,  // immediate
    amqp.Publishing {
      ContentType: "text/plain",
      Body:        []byte(body),
    })
  failOnError(err, "Failed to publish a message")
  log.Printf(" [x] Sent %s\n", body)
}