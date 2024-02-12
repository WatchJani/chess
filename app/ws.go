package app

import (
	"sync"

	"github.com/fasthttp/websocket"
)

type WS struct {
	WSU  websocket.FastHTTPUpgrader
	WSCL *WSCL
}

func NewWS() *WS {
	return &WS{
		WSU: websocket.FastHTTPUpgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		WSCL: NewWSCL(),
	}
}

// web socket client list
type WSCL struct {
	CM             map[string][]*websocket.Conn //for chess party need 2 players
	connectionLost map[string]chan struct{}
	sync.RWMutex
}

func NewWSCL() *WSCL {
	return &WSCL{
		connectionLost: make(map[string]chan struct{}),
		CM:             make(map[string][]*websocket.Conn),
	}
}
