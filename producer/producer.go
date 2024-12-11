package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func produceMessage(message string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		return err
	}
	defer p.Close()

	topic := "test_topic"
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Mensaje producido: %s\n", message)
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		message := r.FormValue("message")
		err := produceMessage(message)
		if err != nil {
			http.Error(w, "Error al enviar el mensaje", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Mensaje enviado: %s", message)
	} else {
		http.Error(w, "Método no soportado", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/send", handler)
	fmt.Println("Productor en ejecución en http://localhost:8080/send")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
