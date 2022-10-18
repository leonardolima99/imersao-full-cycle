package main

import (
	"github.com/codeedu/imersaofsfc2-simulator/infra/kafka"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
/* 	producer := kafka.NewKafkaProducer()
	kafka.Publish("Ola", "route.new-direction", producer)
 */
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
	}
	
}
