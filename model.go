package watchr

import (
	"encoding/json"
	
	"github.com/gorilla/websocket"
)

type connInfo struct {
	conn *websocket.Conn
	room string
	name string
}

type roomInfo struct {
	mediaType string
	media string
}

const (
	opChat       = 0
	opIdentify   = 1
	opUserUpdate = 2
	opRoomUpdate = 3
	opUserSet    = 4
	opRoomSet    = 5
)

type message struct {
	Op   int             `json:"op"`
	Data json.RawMessage `json:"data"`
}

type identify struct {
	Room string `json:"room"`
	Name string `json:"name"`
}

type chat struct {
	Content string `json:"content"`
}

type userset struct {
	Name string `json:"name"`
	Room string `json:"room"`
}

type wmessage struct {
	Op   int         `json:"op"`
	Data interface{} `json:"data"`
}

type wuserupdate struct {
	Name string `json:"name"`
	Room string `json:"room"`
}

const (
	// MediaImage is for static images
	MediaImage = "image"
	// MediaVideo is for a direct video
	MediaVideo = "video"
	// MediaFrame is for a trusted frame (youtube, twitch)
	MediaFrame = "frame"
	// MediaText is for a placeholder text
	MediaText = "text"
)

type wroomupdate struct {
	MediaType string `json:"media_type"`
	Media     string `json:"media"`
}

type wchat struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}
