package handlers

import (
	"bytes"
	"log"

	"go-kafka-basic-consumer/src/interfaces"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandle(t *testing.T) {

	var buf bytes.Buffer
	log.SetOutput(&buf)
	event := interfaces.Event{Key: "testKey", Value: "testValue"}
	Handle(event)

	assert.Contains(t, buf.String(), "Event received - Key: testKey, Value: testValue")
}
