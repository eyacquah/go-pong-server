package websocket

import (
	"encoding/json"
	"log"

	"github.com/eyacquah/react-go-pong-server/pong"
)



func decodeMessage(input string, g *pong.Game) []byte {
	g.Update(input)
	
	// Convert Game to JSON
	str, err := json.Marshal(g)
	if err != nil {
		log.Println(err)
	}

	out := []byte(str)

	return out
}