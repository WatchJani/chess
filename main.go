package main

import "root/app"

func main() {
	srv := app.NewApp()
	ws := app.NewWS()

	srv.CreateWSRoute(ws) //create ws chess route
	srv.CreateHTTPRoute()

	srv.Run(":3000")

	// chess := chess.NewChess()

	// if err := chess.Move(1, 1, 2, 2); err != nil {
	// 	log.Println(err)
	// }

	// chess.Print()
}
