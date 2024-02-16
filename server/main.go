package main

import (
	"root/app"
)

func main() {
	srv := app.NewApp()

	srv.InitHandler()

	srv.Run()
}
