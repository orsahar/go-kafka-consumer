package handlers

import (
	"bytes"
	"log"

	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEvent is a test implementation of the Event interface.
type TestEvent struct {
	key   string
	value string
}

// Key returns the key of the test event.
func (e TestEvent) Key() string {
	return e.key
}

// Value returns the value of the test event.
func (e TestEvent) Value() string {
	return e.value
}

func TestHandle(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil) // Reset log output after the test

	event := TestEvent{key: "testKey", value: "testValue"}
	Handle(event)

	assert.Contains(t, buf.String(), "Event received - Key: testKey, Value: testValue")
}
