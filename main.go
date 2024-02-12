package main

import (
	"fmt"
	"log"
	"root/chess"
)

func main() {
	// srv := app.NewApp()
	// ws := app.NewWS()

	// srv.CreateWSRoute(ws) //create ws chess route
	// srv.CreateHTTPRoute()

	// srv.Run(":3000")

	chess := chess.NewChess()

	if err := chess.Move(1, 1, 2, 1); err != nil {
		log.Println(err)
	}

	fmt.Println(string(chess.Print()))
}
