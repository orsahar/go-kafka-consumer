package interfaces

// Event represents an interface for events with key and value.
type Event interface {
	Key() string
	Value() string
}
