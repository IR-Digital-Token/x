package messages

// Payload is the Message's payload.
type Payload []byte

// Metadata is sent with every message to provide extra context without unmarshaling the message payload.
type Metadata map[string]string

// Message is the basic transfer unit.
// Messages are emitted by Publishers and received by Subscribers.
type Message struct {
	// UUID is an unique identifier of message.
	UUID string

	// Metadata contains the message metadata.
	Metadata Metadata

	// Payload is the message's payload.
	Payload Payload
}

// Get returns the metadata value for the given key. If the key is not found, an empty string is returned.
func (m Metadata) Get(key string) string {
	if v, ok := m[key]; ok {
		return v
	}

	return ""
}

// Set sets the metadata key to value.
func (m Metadata) Set(key, value string) {
	m[key] = value
}

// NewMessage creates a new Message with given uuid and payload.
func NewMessage(payload Payload) *Message {
	return &Message{
		Metadata: make(map[string]string),
		Payload:  payload,
	}
}
