package srv

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var connections = make(chan *websocket.Conn)

func handleWebSocket(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		connections <- c

		defer func() {
			c.Close()
			delete(connections, c)
		}()

		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				fmt.Printf("read: %s\n", err)
			}
			fmt.Println("recv:", message)
			for conn := range connections {
				err = conn.WriteMessage(mt, message)
				if err != nil {
					fmt.Println("write:", err)
					break
				}
			}
		}
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
