package server

import (
	"log"
	"net/http"

	"github.com/mitchellh/mapstructure"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a client connected to the server
type Client struct {
	conn  *websocket.Conn
	write chan Message
}

func (c *Client) readLoop() {
	defer c.conn.Close()

	for {
		var msg Message
		if err := c.conn.ReadJSON(&msg); err != nil {
			c.writeError(&err)
			return
		}

		switch msg.OpCode {
		case OpIdentify:
			var i Identify
			if err := mapstructure.Decode(msg.Payload, &i); err != nil {
				c.writeError(&err)
				return
			}
			// TODO: multiple channels lol
			j := Message{
				OpCode: OpChangeMedia,
				Payload: ChangeMedia{
					MediaURL: "https://www.youtube.com/watch?v=S-AdFbElLts",
				},
			}
			c.write <- j
		default:
			log.Println("(hey) unknown msg", msg)
		}
	}
}

func (c *Client) writeLoop() {
	defer c.conn.Close()

	for {
		select {
		case msg, ok := <-c.write:
			if !ok {
				c.writeError(nil)
				return
			}

			if err := c.conn.WriteJSON(msg); err != nil {
				c.writeError(&err)
				return
			}
		}
	}
}

func (c *Client) writeError(e *error) {
	var text string
	if e != nil {
		text = (*e).Error()
	} else {
		text = "unspecified error"
	}
	log.Println("(err)", text)

	close := websocket.FormatCloseMessage(websocket.CloseInternalServerErr, text)
	c.conn.WriteMessage(websocket.CloseMessage, close)
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	client := Client{
		conn:  conn,
		write: make(chan Message, 128),
	}

	go client.writeLoop()
	go client.readLoop()

	hello := Message{
		OpCode:  OpHello,
		Payload: 0,
	}
	client.write <- hello
}
