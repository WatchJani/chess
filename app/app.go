package app

import (
	"log"
	"time"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type App struct {
	Router *router.Router
}

func NewApp() *App {
	return &App{
		Router: router.New(),
	}
}

func (a *App) CreateWSRoute(handler *WS) {
	a.Router.Handle("GET", "/create", handler.CreateGame)
	a.Router.Handle("GET", "/join", handler.JoinGame)
}

func (a *App) CreateHTTPRoute() {
	a.Router.Handle("GET", "/test", TestPlay)
}

func (a *App) Run(port string) {
	srv := &fasthttp.Server{
		Handler:          a.Router.Handler,
		DisableKeepalive: false,
		MaxConnsPerIP:    8,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
		WriteTimeout:     15 * time.Second,
		ReadTimeout:      15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe(port))
}
