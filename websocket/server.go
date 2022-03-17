package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eyacquah/react-go-pong-server/pong"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func reader(conn *websocket.Conn) {
	g := &pong.Game{}

	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		input := string(p)
		// log.Println("From Client",string(p))

		msg := decodeMessage(input, g)

		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return 
		}
	}
}




func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {return true}

	ws, err := upgrader.Upgrade(w, r, nil)
	
	if err != nil {
		log.Println(err)
	}

	log.Println("Client connected successfully")
	reader(ws)
}

func SetupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)
}