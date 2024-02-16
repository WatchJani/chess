package handler

import (
	"errors"
	"log"
	"math"
	"math/rand"
	"sync"

	"github.com/fasthttp/websocket"
)

type WS struct {
	WSU         websocket.FastHTTPUpgrader
	GameSession map[string]*GameSession
	sync.RWMutex
}

func NewWS() WS {
	return WS{
		WSU: websocket.FastHTTPUpgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		GameSession: make(map[string]*GameSession),
	}
}

func (ws *WS) CreateGameSession(host string) *GameSession {
	ws.Lock()
	ws.GameSession[host] = NewGameSession(host)
	ws.Unlock()
	return ws.GameSession[host]
}

func SendMessage(conn *websocket.Conn, message string) {
	conn.WriteMessage(websocket.TextMessage, []byte(message))
}

func (ws *WS) CheckJoin(host string) error {
	if _, ok := ws.GameSession[host]; !ok {
		return errors.New("this game not exist")
	}

	return nil
}

func (ws *WS) GetSession(host string) *GameSession {
	return ws.GameSession[host]
}

func (ws *WS) SendMove(host string, msg []byte) {
	for _, conn := range ws.GameSession[host].Room {
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func (ws *WS) ErrorHandler(mt int, err error, host string) bool {
	if mt == websocket.CloseMessage {
		log.Println("Client [" + host + "] closed connection | Message: CloseMessage")
		// break close
		return true
	}
	if mt == websocket.CloseAbnormalClosure {
		log.Println("Client [" + host + "] closed connection | Message: CloseAbnormalClosure")
		// break close
		return true
	}
	if err != nil {
		if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
			log.Println("Client ["+host+"] closed connection | Message:", mt)
			// break close
			return true
		}
	}

	return false
}

type GameSession struct {
	GameID int
	Host   string
	Close  chan struct{}
	Room   []*websocket.Conn
	GameLoop
}

func NewGameSession(host string) *GameSession {
	return &GameSession{
		Host:     host,
		GameID:   rand.Intn(math.MaxInt),
		Close:    make(chan struct{}),
		Room:     make([]*websocket.Conn, 0, 2),
		GameLoop: NewGameLoop(),
	}
}

func (g *GameSession) CloseConn() {
	<-g.Close
}

func (g *GameSession) Check() error {
	sessionPlayer := len(g.Room)

	if g.Checker == sessionPlayer {
		g.Left = false
		return errors.New("player is left")
	}

	g.Checker = sessionPlayer
	return nil
}

func (g GameSession) GetPlayer() *websocket.Conn {
	return g.Room[g.Counter%len(g.Room)]
}

func (g *GameSession) CloseSession() {
	if g.Left {
		g.Close <- struct{}{}
	}
}

func (g *GameSession) HostMessage(msg string) {
	g.Room[0].WriteMessage(websocket.TextMessage, []byte(msg))
}

func (g *GameSession) JoinMessage(msg string) {
	g.Room[1].WriteMessage(websocket.TextMessage, []byte(msg))
}

func (g *GameSession) SessionAppend(c *websocket.Conn) error {
	if len(g.Room) > 2 {
		return errors.New("room is full")
	}

	g.Room = append(g.Room, c)

	return nil
}

type GameLoop struct {
	Counter int
	Checker int
	Left    bool
}

func NewGameLoop() GameLoop {
	return GameLoop{
		Checker: -1,
		Left:    true,
	}
}

func (g *GameLoop) NextMove() {
	g.Counter++
}

func (g *GameLoop) IsGameEnd() bool {
	return g.Counter == 100
}
