package main

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"

	"github.com/anthdm/hollywood/actor"
	"github.com/gorilla/websocket"
)

type PlayerSession struct {
	sessionID int
	clientID  int
	username  string
	inLobby   bool
	conn      *websocket.Conn
}

func newPlayerSession(sid int, conn *websocket.Conn) actor.Producer {
	return func() actor.Receiver {
		return &PlayerSession{
			conn:      conn,
			sessionID: sid,
		}
	}
}

func (s *PlayerSession) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.readLoop()
	}
}

func (s *PlayerSession) readLoop {

	for {
		if err :=	s.conn.ReadJSON(msg); err != nil {
			fmt.Println("read error", err)
			return
		}
	}
}

type GameServer struct {
	ctx      *actor.Context
	sessions map[*actor.PID]struct{}
}

func newGameServer() actor.Receiver {
	return &GameServer{
		sessions: make(map[*actor.PID]struct{}),
	}
}

func (s *GameServer) Receive(c *actor.Context) {
	switch msg := c.Message().(type) {
	case actor.Started:
		s.startHTTP()
		s.ctx = c
		_ = msg
	}
}

func (s *GameServer) startHTTP() {
	fmt.Println("starting HTTP server on port: 40000")
	go func() {
		http.HandleFunc("/ws", s.handleWS)
		http.ListenAndServe(":40000", nil)
	}()
}

// handles upgrades of the WS
func (s *GameServer) handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		fmt.Println("ws upgrade err:", err)
		return
	}
	fmt.Println("new client trying to connect")
	sid := rand.Intn(math.MaxInt)
	pid := s.ctx.SpawnChild(newPlayerSession(sid, conn), fmt.Sprintf("session_%d", sid))
	s.sessions[pid] = struct{}{}
	fmt.Printf("client with sid %d and pid %s just connected\n", sid, pid)
}

func main() {
	fmt.Println("error with actor engine")
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
	}
	e.Spawn(newGameServer, "server")
	select {}
}
