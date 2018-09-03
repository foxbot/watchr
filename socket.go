package watchr

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var gatewayUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: security alert
	},
}

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

	name := i.Name
	if name == "" {
		i := rand.Intn(32678)
		name = fmt.Sprintf("anon%d", i)

		uu := wmessage{
			Op: opUserUpdate,
			Data: wuserupdate{
				Name: name,
			},
		}
		err = c.WriteJSON(uu)
		if err != nil {
			w.errors <- err
		}
	}

	ci := &connInfo{
		conn: c,
		room: i.Room,
		name: name,
	}
	// TODO: more unique id
	log.Println("new client from", c.RemoteAddr().String())
	w.clients[c.RemoteAddr().String()] = ci
	c.SetCloseHandler(func(code int, text string) error {
		delete(w.clients, c.RemoteAddr().String())
		return nil
	})

	w.sendRoomInfo(c, ci.room)

	for {
		var m message
		err := c.ReadJSON(&m)
		if err != nil {
			w.errors <- err
			continue
		}
		switch m.Op {
		case opChat:
			var d chat
			err = json.Unmarshal(m.Data, &d)
			if err != nil {
				w.errors <- err
				break
			}
			wm := wmessage{
				Op: opChat,
				Data: wchat{
					Author:  ci.name,
					Content: d.Content,
				},
			}
			j, err := json.Marshal(wm)
			if err != nil {
				w.errors <- err
				break
			}
			w.broadcast(ci.room, j)
		case opUserSet:
			var d userset
			err = json.Unmarshal(m.Data, &d)
			if err != nil {
				w.errors <- err
				break
			}
			var wd wuserupdate
			if d.Name != "" {
				// TODO: name validation?
				ci.name = d.Name
				wd.Name = d.Name
			}
			if d.Room != "" {
				// TODO: room validation/passwording
				ci.room = d.Room
				wd.Room = d.Room
				w.sendRoomInfo(c, d.Room)
			}
			wm := wmessage{
				Op:   opUserUpdate,
				Data: wd,
			}
			err = ci.conn.WriteJSON(wm)
			if err != nil {
				w.errors <- err
			}
		}
	}
}

func (w *Watchr) broadcast(room string, data []byte) error {
	t := make([]*websocket.Conn, 0)
	for _, ci := range w.clients {
		if ci.room == room {
			t = append(t, ci.conn)
		}
	}
	m, err := websocket.NewPreparedMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}
	log.Println("broadcasting to", len(t), "clients in", room)
	for _, c := range t {
		err = c.WritePreparedMessage(m)
		if err != nil {
			w.errors <- err
			continue
		}
	}
	return nil
}

func (w *Watchr) sendRoomInfo(c *websocket.Conn, r string) {
	room, ok := w.rooms[r]
	if !ok {
		room = &roomInfo{
			mediaType: MediaText,
			media:     "this is a new room named " + r,
		}
		w.rooms[r] = room
	}

	ru := wmessage{
		Op: opRoomUpdate,
		Data: wroomupdate{
			MediaType: room.mediaType,
			Media:     room.media,
		},
	}
	err := c.WriteJSON(ru)
	if err != nil {
		w.errors <- err
	}
}
