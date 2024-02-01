//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial("amqp://WRC:Wrc@123@localhost:5672/shared")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"hello:exchange",
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a exchange")

	// 声明一个消息队列
	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	ch.QueueBind(
		q.Name,
		q.Name,
		"hello:exchange",
		false,
		nil,
	)

	// 循环发送消息
	for i := 0; i < 100; i++ {
		// 发布消息
		body := fmt.Sprintf("Hello World! * %v", i+1)
		err = ch.Publish(
			"hello:exchange", // exchange
			q.Name,           // routing key
			false,            // mandatory
			false,            // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		log.Printf(" [x] Sent %s", body)
		failOnError(err, "Failed to publish a message")

		time.Sleep(1 * time.Second)
	}
}
