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
func InfoOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s", err)
	} else {
		log.Printf("%s", msg)
	}
}

func main() {
	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	queueNum, err := ch.QueueDelete("hello", true, true, false)
	InfoOnError(err, fmt.Sprintf("Close queue number %v", queueNum))

	time.Sleep(10 * time.Second)

	err = ch.ExchangeDelete("hello:exchange", true, false)
	InfoOnError(err, fmt.Sprintf("Close Exchange %v", "hello:exchange"))

	// 等待程序结束
	log.Printf(" [*] Close exchange and queue")
}
