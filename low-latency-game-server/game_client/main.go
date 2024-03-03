package main

import (
	"log"
	"math"
	"math/rand"

	"github.com/clgerhardt/gameserver/types"
	"github.com/gorilla/websocket"
)

const wsServerEndpoint = "ws://localhost:40000/ws"

type GameClient struct {
	conn     *websocket.Conn
	clientId int
	username string
}

func newGameClient(conn *websocket.Conn, username string) *GameClient {
	return &GameClient{
		clientId: rand.Intn(math.MaxInt),
		username: username,
		conn:     conn,
	}
}

func (c *GameClient) login() error {
	return c.conn.WriteJSON(types.Login{
		ClientId: c.clientId,
		Username: c.username,
	})
}

func main() {

	dailer := websocket.Dialer{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	conn, _, err := dailer.Dial(wsServerEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	c := newGameClient(conn, "Christian")
	if err := c.login(); err != nil {
		log.Fatal(err)
	}

}
