package websocket

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/eyacquah/react-go-pong-server/pong"
)

func getJSON(g *pong.Game) []byte {
	str, err := json.Marshal(g)
	if err != nil {
		log.Println(err)
	}

	return str
}


func decodeMessage(input string, g *pong.Game) []byte {
	
	switch input {
	case "INIT":
		g.Init()

	case "UPDATE":
		g.UpdateBall()

	default:
		if strings.Contains(input, "P1") || strings.Contains(input, "P2") {
			g.UpdatePads(input)
			g.UpdateBall()
		}
	}
	str := getJSON(g)
	out := []byte(str)

	return out
}