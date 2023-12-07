package request

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type Param struct {
	Url       string
	Body      interface{}
	Method    string
	Headers   map[string]interface{}
	BodyType  string
	CreateLog bool
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},

	ReadBufferSize:  16384, // 16 KB
	WriteBufferSize: 16384, // 16 KB
}
