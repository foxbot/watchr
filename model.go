package watchr

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type connInfo struct {
	conn    *websocket.Conn
	channel string
}

const (
	opHello    = 0
	opIdentify = 1
)

type message struct {
	Op   int             `json:"op"`
	Data json.RawMessage `json:"data"`
}
type wmessage struct {
	Op   int         `json:"op"`
	Data interface{} `json:"data"`
}

type hello struct {
	ID string `json:"id"`
}

type identify struct {
	Channel string `json:"channel"`
}
