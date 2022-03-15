package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/eyacquah/react-go-pong-server/pong"

)

type Message struct {
	pong.Ball
	pong.Paddle
	Type        	string 		`json:"type"`
	Error 			string		`json:"error"`
	ScreenWidth  	int    		`json:"screenWidth"`
	ScreenHeight 	int    		`json:"screenHeight"`

}

func getJSON(msg Message) []byte {
	str, err := json.Marshal(msg)

	if err != nil {
		log.Println(err)
	}

	return str
}

func parseJSON(in []byte ,out Message) Message {
	err := json.Unmarshal([]byte(in), &out)

		if err != nil {
			log.Println(err)
			// return
		}

	return out
}

func decodeMessage(in []byte) []byte {
	var msg Message

	message := parseJSON(in, msg)
	var ball = pong.Ball{}
	var paddle = pong.Paddle{}

	switch message.Type {
	case "INIT":
		message.Type = "INIT"
		message.ScreenHeight = 600
		message.ScreenWidth = 600

		paddle.Height = pong.InitPaddleHeight
		paddle.Width = pong.InitPaddleWidth
		message.Paddle = paddle

		ball.Radius = pong.InitBallRadius
		ball.Position = pong.Position{
			X: 10,
			Y: 10,
		}
		message.Ball = ball

	default:
		message.Type = "ERROR"
		message.Error = "Something Went Wrong :("
	}

	fmt.Println("MSG",message)

	str := getJSON(message)
	
	out := []byte(str)
	fmt.Println("OUT-STR", string(out))
	

	return out
}