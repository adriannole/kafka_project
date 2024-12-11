package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "test_group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Fatalf("Error al crear el consumidor: %s", err)
	}
	defer c.Close()

	c.SubscribeTopics([]string{"test_topic"}, nil)

	fmt.Println("Consumidor escuchando mensajes...")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Mensaje recibido: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Error al leer el mensaje: %v\n", err)
		}
	}
}
