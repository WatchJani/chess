package app

import (
	"fmt"
	"log"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

func TestPlay(ctx *fasthttp.RequestCtx) {
	fmt.Println(string(ctx.Host()))
}

func (w *WS) CreateGame(ctx *fasthttp.RequestCtx) {
	err := w.WSU.Upgrade(ctx, func(c *websocket.Conn) {
		defer func(ws *websocket.Conn) {
			err := ws.Close()
			if err != nil {
				log.Println(err)
			}
		}(c)

		w.WSCL.Lock()
		host := ctx.RemoteAddr().String()
		fmt.Println(host)
		w.WSCL.CM[host] = append(w.WSCL.CM[host], c)
		w.WSCL.Unlock()

		c.WriteMessage(websocket.TextMessage, []byte("Game is created")) //witch messageType
		go w.CreateNewGame(host)

		w.WSCL.connectionLost[host] = make(chan struct{})

		fmt.Println("conn host", host)
		<-w.WSCL.connectionLost[host] //witch game we need to wait

		fmt.Println("game host is left")
	})

	if err != nil {
		fmt.Println("WS Handshake Error!")
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println(err)
		}
		return
	}
}

func (w *WS) JoinGame(ctx *fasthttp.RequestCtx) {
	err := w.WSU.Upgrade(ctx, func(c *websocket.Conn) {
		defer func(wc *websocket.Conn) {
			if err := wc.Close(); err != nil {
				log.Println(err)
			}
		}(c)

		w.WSCL.RLock()
		host := string(ctx.Request.Header.Peek("join"))
		fmt.Println(host)
		list, ok := w.WSCL.CM[host]
		w.WSCL.RUnlock()

		//error handle
		if !ok {
			return
		}

		if len(list) > 2 {
			return
		}

		//add connection
		fmt.Println("user connected")
		list = append(list, c)

		list[0].WriteMessage(websocket.TextMessage, []byte(ctx.RemoteAddr().String())) //witch messageType
		list[1].WriteMessage(websocket.TextMessage, []byte("join the game"))

		<-w.WSCL.connectionLost[host]
		fmt.Println("player joined left")
	})

	if err != nil {
		fmt.Println("WS Handshake Error!")
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println(err)
		}
		return
	}
}

// func (w *WS) CreateGame(ctx *fasthttp.RequestCtx) {
// 	//some http logic
// 	//i dont needed in this case

// 	if err := w.WSU.Upgrade(ctx, w.Chess); err != nil {
// 		fmt.Println("WS Handshake Error!")
// 		if _, ok := err.(websocket.HandshakeError); ok {
// 			log.Println(err)
// 		}
// 		return
// 	}
// }

// //Radi :D

// // go routine
// func (w *WS) Chess(ws *websocket.Conn) {
// 	// w.WSCL.Lock()

// 	// w.WSCL.Unlock()

// 	//func is because we have more then one line
// 	//every connection need to close
// 	defer func(ws *websocket.Conn) {
// 		err := ws.Close()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}(ws)

// 	go w.CreateNewGame()

// 	//we need to wait to game end or players left
// 	<-w.WSCL.connectionLost

// 	//defer will be executed
// }

func (w *WS) CreateNewGame(host string) {
	time.Sleep(15 * time.Second)
	fmt.Println("game is initialized")

	fmt.Println("game host", host)
	for index := 0; index < 2; index++ {
		w.WSCL.connectionLost[host] <- struct{}{}
	}
}
