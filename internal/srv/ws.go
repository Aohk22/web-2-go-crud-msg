package srv

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan []byte // the hub sends broadcast messages to this channel
}

type Hub struct {
	clients map[*Client]bool
	reg chan *Client
	unreg chan *Client
	broadcast chan []byte
}

func (c *Client) listenAndWriteToHub() {
	defer func() {
		c.hub.unreg <- c
		c.conn.Close()
	}()
	// not set read deadline yet
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.hub.broadcast <- message
	}
}

func (c *Client) listenForHubBroadcast() {
	defer func() {
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <- c.send: // c.send is accessed by hub to write broadcast messages
			// not set write deadline
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			log.Println("got broadcast message from hub")

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			log.Println("wrote message to conn")

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write([]byte("\n"))
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.reg:
			h.clients[client] = true
			log.Println("client registered")
		case client := <-h.unreg:
			log.Println("client unregistering")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			log.Println("got message from client, broadcasting")
			for c := range h.clients {
				select {
				case c.send <- message:
				default:
					close(c.send)
					delete(h.clients, c)
				}
			}
		}
	}
}

func NewHub() (*Hub) {
	return &Hub{
		clients: make(map[*Client]bool),
		reg: make(chan *Client),
		unreg: make(chan *Client),
		broadcast: make(chan []byte),
	}
}

var upgrader = websocket.Upgrader{}

func handleWebSocket(ctx context.Context, hub *Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		client := &Client{ hub: hub, conn: c, send: make(chan []byte, 256) }
		log.Println("queueing client for hub register")
		client.hub.reg <- client

		log.Println("starting listenAndWriteToHub()")
		go client.listenAndWriteToHub()
		log.Println("starting listenForHubBroadcast()")
		go client.listenForHubBroadcast()
	}
}

		// defer c.Close()
		// for {
		// 	mt, message, err := c.ReadMessage()
		// 	if err != nil {
		// 		fmt.Println("read:", err)
		// 		break
		// 	}
		// 	fmt.Println("recv:", message)
		// 	err = c.WriteMessage(mt, message)
		// 	if err != nil {
		// 		fmt.Println("write:", message)
		// 		break
		// 	}
		// }
