package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"

	"github.com/fasthttp/websocket"
)

const Endpoint string = "ws://localhost:8800/create"

const (
	ErrBadHandshake = "Error on connecting to Websocket Server: Bad Handshake | Check your login info and be sure that user isn't already logged in"
	ErrConnection   = "Error connecting to Websocket Server | Error: "
)

type Client struct {
	ClientID   int
	ClientName string
	Conn       *websocket.Conn
}

func NewClient(clientName string, conn *websocket.Conn) *Client {
	return &Client{
		ClientID:   rand.Intn(math.MaxInt),
		ClientName: clientName,
		Conn:       conn,
	}
}

func (c *Client) Join(msg []byte) error {
	return c.Conn.WriteMessage(websocket.TextMessage, msg)
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(Endpoint, nil)
	if err != nil {
		if err == websocket.ErrBadHandshake {
			fmt.Println(ErrBadHandshake)
			os.Exit(0)
		} else if websocket.IsUnexpectedCloseError(err) {
			os.Exit(0)
		} else {
			fmt.Println(ErrConnection, err.Error())
			os.Exit(0)
		}
	}

	fmt.Println("Client has been connected!")
	defer func(c *websocket.Conn) {
		err = c.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	c := NewClient("John", conn)
	//make move
	c.Join([]byte("A2 A3"))
}
