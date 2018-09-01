package watchr

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var gatewayUpgrader = websocket.Upgrader{}

func (w *Watchr) onGateway(rw http.ResponseWriter, r *http.Request) {
	conn, err := gatewayUpgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(500)
		return
	}

	go w.read(conn)
}

func (w *Watchr) read(c *websocket.Conn) {
	defer func() {
		c.Close()
	}()

	// read ident
	var m message
	err := c.ReadJSON(&m)
	if err != nil {
		w.errors <- err
		return
	}
	if m.Op != opIdentify {
		c.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(
			4000,
			"expected identify!"),
			time.Time{})
		c.Close()
	}
	var i identify
	err = json.Unmarshal(m.Data, &i)
	if err != nil {
		w.errors <- err
		return
	}

	ci := connInfo{
		conn:    c,
		channel: i.Channel,
	}
	w.clients["todo"] = ci

	for {
		var m message
		err := c.ReadJSON(&m)
		if err != nil {
			w.errors <- err
			continue
		}

	}
}
