package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eyacquah/react-go-pong-server/websocket"
)


const (
	windowWidth = 600
	windowHeight = 600
)

func main() {
	fmt.Println("React Go Pong Server Is Live!")
	
	websocket.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))

}
