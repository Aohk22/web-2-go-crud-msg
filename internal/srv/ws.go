package srv

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
	stores *Stores
}

func (c *Client) listenAndWriteToHub() {
	defer func() {
		c.hub.unreg <- c
		c.conn.Close()
	}()
	// TODO: not set read deadline yet
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Printf("%v\n", err)
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
			log.Println("got message from client, broadcasting and saving messeage to database")
			d := struct { Time, Content, UserId, RoomId string}{}
			err := json.Unmarshal(message, &d)
			if err != nil {
				log.Println("socket: failed to unmarshal message")
				log.Println(err)
				return
			}
			log.Println(d)
			log.Println(d.Time)
			i, err := strconv.ParseInt(d.Time, 10, 64)
			if err != nil {
				log.Println("socket: failed to parse int time")
				return
			}
			tm := time.Unix(i, 0).UTC().Format("2006-01-02 15:04:05-07") // 0 nsec
			uidInt, err := strconv.ParseInt(d.UserId, 10, 32)
			ridInt, err := strconv.ParseInt(d.RoomId, 10, 32)
			if err != nil { 
				log.Println("socket: failed to parse ids")
				log.Println(err)
			}
			_, err = h.stores.MessageStore.AddMessage(context.Background(), tm, d.Content, uint32(uidInt), uint32(ridInt)) // TODO potential bug here
			if err != nil {
				log.Println("socket: failed to save to database therefore not putting to echoing to clients")
				log.Println(err)
				return
			}
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

func NewHub(stores *Stores) (*Hub) {
	return &Hub{
		clients: make(map[*Client]bool),
		reg: make(chan *Client),
		unreg: make(chan *Client),
		broadcast: make(chan []byte),
		stores: stores,
	}
}

// TODO: FIX THIs
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	Subprotocols: []string{"wamp"}, // for token smuggling, not really needed now
}

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
