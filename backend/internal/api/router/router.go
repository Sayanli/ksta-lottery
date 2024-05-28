package router

import (
	"log"
	"lottery/internal/api/handler"
	"lottery/internal/service"
	"lottery/pkg/websocket"
	"net/http"
)

type Server struct {
	hander *handler.Handler
}

func NewServer(service *service.Service) *Server {
	return &Server{
		hander: handler.NewHandler(service),
	}
}

func (s *Server) serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err.Error())
	}

	client := &websocket.Client{
		Handler: s.hander,
		Conn:    conn,
		Pool:    pool,
	}

	pool.Register <- client
	client.Read(s.hander)
}

func (s *Server) SetupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		s.serveWs(pool, w, r)
	})
}
