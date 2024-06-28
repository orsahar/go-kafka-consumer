package handlers

import (
	"go-kafka-basic-consumer/src/interfaces"
	"log"
)

// Handle logs the event to the console.

// Handle logs the event to the console.
func Handle(event interfaces.Event) {
	log.Printf("Event received - Key: %s, Value: %s\n", event.Key(), event.Value())
}
