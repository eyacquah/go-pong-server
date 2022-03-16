package pong

import "strings"

type Paddle struct {
	Position
	Score  int			`json:"score"`
	Speed  float32		`json:"speed"`
	Width  int			`json:"width"`
	Height int			`json:"height"`
}

const (
	InitPaddleWidth  = 15
	InitPaddleHeight = 70
	InitPaddleShift = 20
)

func (p *Paddle) Update(input string) {

	if strings.Contains(input, "UP_PRESS") {
		// Keep moving up
		p.Y -= p.Speed
	} else if strings.Contains(input, "DOWN_PRESS") {
		// Keep moving down
		p.Y += p.Speed
	} else {
		// Key was released. Stop moving
		p.Y += 0
	}

	// Detect walls.
	// Top wall
	if p.Y < 0 {
		p.Y = 0
	}

	// Bottom wall
	if p.Y + float32(p.Height) > ScreenHeight {
		p.Y = float32(ScreenHeight - p.Height)
	}
}