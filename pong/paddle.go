package pong

type Paddle struct {
	Position
	Score int
	Speed float32
	Width int
	Height int
}

const (
	InitPaddleWidth = 20
	InitPaddleHeight = 70
)