package handler

import (
	"fmt"
	"log"
	"root/chess"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

const (
	ErrWrongMoveInput = "wrong move or input:"
	HostConnected     = "Host is Connected!"
	YourTurn          = "Your Turn: "
)

func (ws *WS) Create(ctx *fasthttp.RequestCtx) {
	if err := ws.WSU.Upgrade(ctx, ws.CreateGame); err != nil {
		log.Println(err)
	}
}

func (ws *WS) Join(ctx *fasthttp.RequestCtx) {
	host := ctx.Request.Header.Peek("host")

	if err := ws.CheckJoin(string(host)); err != nil {
		log.Println(err)
		return
	}

	if err := ws.WSU.Upgrade(ctx, func(c *websocket.Conn) {
		ws.JoinGame(c, string(host))
	}); err != nil {
		log.Println(err)
	}
}

func Close(c *websocket.Conn) {
	if err := c.Close(); err != nil {
		log.Println(err)
	}
}

func (ws *WS) JoinGame(c *websocket.Conn, host string) {
	defer Close(c)

	fmt.Println("Client is connected!")
	session := ws.GetSession(host)

	//subscribe
	session.SessionAppend(c)

	//checkpoint
	session.CloseConn()

	//delete from list

	session.JoinMessage("Host is left, game is end :D")
}

// host create new go routine
func (ws *WS) CreateGame(c *websocket.Conn) {
	//Close connection
	defer Close(c)

	fmt.Println(HostConnected)

	//get host ip address
	host := c.RemoteAddr().String()
	fmt.Println(host)

	//add host to session game
	session := ws.CreateGameSession(host)

	//subscribe to game
	session.SessionAppend(c)

	//create new game instance
	game := chess.NewChess()
	fmt.Println("game is initialized")

	//game loop
game:
	for {
		//Check if all player is here
		if err := session.Check(); err != nil {
			log.Println(err)
			break
		}

		//get player for next move
		playerTurn := session.GetPlayer()

		SendMessage(playerTurn, YourTurn)

		for {
			mt, msg, err := playerTurn.ReadMessage()

			//conn error handler
			if ws.ErrorHandler(mt, err, host) {
				break game
			}

			//this is important part of code, but that cant be here
			if len(msg) != 5 {
				SendMessage(playerTurn, ErrWrongMoveInput)
				continue
			}

			//create move in chess
			if err := game.Move(chess.Parse(msg)); err != nil {
				SendMessage(playerTurn, ErrWrongMoveInput)
				continue
			}

			break
		}

		//check if is end of game
		if session.IsGameEnd() {
			break
		}

		//send move to both player
		ws.SendMove(host, game.Print())

		//next move
		session.NextMove()
	}

	//Close session
	session.CloseSession()
}
