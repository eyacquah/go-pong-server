package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eyacquah/react-go-pong-server/websocket"
)

func main() {
	
	// [START setting_port]
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	
	websocket.SetupRoutes()
	
	log.Printf("React Go Pong Server Is Live on port %s!", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
