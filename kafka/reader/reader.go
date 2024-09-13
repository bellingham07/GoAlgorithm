package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{"localhost:9092"},
		Topic:          "cao",
		CommitInterval: 1 * time.Second, // 隔一段时间提交offset
		//GroupID:        "zero_chat",
		StartOffset: kafka.FirstOffset, // 新的consumer从最开始的，消息开始消费，仅对刚开始有效
	})
	for {
		fmt.Println("start listening...")
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("reader err:", err)
		}
		fmt.Println("msg value:", string(msg.Value))
	}
}
