package server

const (
	// OpHello s->c
	OpHello int = 0
	// OpIdentify c->s
	OpIdentify = 1
	// OpChangeMedia s->c
	OpChangeMedia = 2
	// OpChat s->c
	OpChat = 3
	// OpWrite c->s
	OpWrite = 4
)

// Message is any incoming or outgoing message
type Message struct {
	OpCode  int         `json:"op"`
	Payload interface{} `json:"d"`
}

// Identify is an incoming payload containing information about the connection
type Identify struct {
	Channel string `json:"channel"`
}

// ChangeMedia is an outgoing payload instructing the client to change its media
type ChangeMedia struct {
	MediaURL string `json:"media_url"`
}
