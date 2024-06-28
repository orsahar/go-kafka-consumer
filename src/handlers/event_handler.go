package handlers

import (
	"fmt"
	"go-kafka-basic-consumer/src/interfaces"
)

// Handle logs the event to the console.
func Handle(event interfaces.Event) {
	fmt.Printf("Event received - Key: %s, Value: %s\n", event.Key, event.Value)
}
