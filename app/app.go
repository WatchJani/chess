package app

import (
	"log"
	"root/app/handler"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type App struct {
	Router *router.Router
	handler.WS
}

func NewApp() *App {
	return &App{
		Router: router.New(),
		WS:     handler.NewWS(),
	}
}

func (a *App) InitHandler() {
	a.Router.Handle("GET", "/create", a.Create)
	a.Router.Handle("GET", "/join", a.Join)
}

func (a *App) Run() {
	srv := &fasthttp.Server{
		Handler:          a.Router.Handler,
		DisableKeepalive: false,
		MaxConnsPerIP:    8,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
		WriteTimeout:     15 * time.Second,
		ReadTimeout:      15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe("localhost:8800"))
}
