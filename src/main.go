package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go-kafka-basic-consumer/src/handlers"
	"go-kafka-basic-consumer/src/interfaces"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
)

// KafkaEventConsumer is responsible for consuming events from Kafka.
type KafkaEventConsumer struct {
	consumer *kafka.Consumer
}

// NewKafkaEventConsumer creates a new KafkaEventConsumer.
func NewKafkaEventConsumer(brokers, groupID, topic string) (*KafkaEventConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	err = c.Subscribe(topic, nil)
	if err != nil {
		return nil, err
	}

	return &KafkaEventConsumer{consumer: c}, nil
}

// Consume starts consuming events from Kafka and logs them to the console.
func (k *KafkaEventConsumer) Consume() {
	for {
		msg, err := k.consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Error reading message: %v\n", err)
			continue
		}

		event := interfaces.Event{
			Key:   string(msg.Key),
			Value: string(msg.Value),
		}

		handlers.Handle(event)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		os.Exit(1)
	}

	brokers := os.Getenv("BROKERS")
	groupID := os.Getenv("GROUP_ID")
	topic := os.Getenv("TOPIC")

	// Initialize the Kafka consumer.
	kafkaConsumer, err := NewKafkaEventConsumer(brokers, groupID, topic)
	if err != nil {
		fmt.Printf("Error creating Kafka consumer: %v\n", err)
		os.Exit(1)
	}

	// Handle graceful shutdown.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go kafkaConsumer.Consume()

	<-sigs
	fmt.Println("Shutting down...")
	kafkaConsumer.consumer.Close()
}
