package pong

import (
	"fmt"
	"strings"
)

type Game struct {
	Ball		*Ball		`json:"ball"`
	Player1		*Paddle		`json:"player1"`
	Player2		*Paddle		`json:"player2"`
	MaxScore	int			`json:"maxScore"`
}

const (
	initBallVelocity = 5.0
	initPaddleSpeed  = 10.0



	ScreenWidth  = 900
	ScreenHeight = 500

	InitMaxScore = 10
)

func (g *Game) Init() {
	fmt.Println("INITIALIZED GAME!!")
		g.MaxScore = 10

	g.Player1 = &Paddle{
		Position: Position{
			X: InitPaddleShift,
			Y: ScreenHeight / 2 - 35,
		},
		Score:  0,
		Speed:  initPaddleSpeed,
		Width:  InitPaddleWidth,
		Height: InitPaddleHeight,
		HasWon: false,
	}

	g.Player2 = &Paddle{
		Position: Position{
			X: ScreenWidth - ( InitPaddleShift * 2),
			Y: ScreenHeight / 2 - 35,
		},
		Score:  0,
		Speed:  initPaddleSpeed,
		Width:  InitPaddleWidth,
		Height: InitPaddleHeight,
		HasWon: false,
	}

	g.Ball = &Ball{
		Position: Position{
			X: InitPaddleShift + InitBallRadius,
			Y: ScreenHeight / 2,
		},

		Radius:    InitBallRadius,
		XVelocity: initBallVelocity,
		YVelocity: initBallVelocity,
	}
}

func (g *Game) UpdatePads(input string) {
	if strings.Contains(input, "P1") {
		g.Player1.Update(input)
	} else if strings.Contains(input, "P2") {
		g.Player2.Update(input)
	}
}


// Inputs
// INIT, UPDATE, RESET

func (g *Game) Update(input string) {
	switch input {
	case "INIT":
		g.Init()
	
	case "UPDATE":
		g.Ball.Update(g.Player1, g.Player2)

		if g.Player1.Score >= InitMaxScore {
			g.Player1.HasWon = true
			
		} else if g.Player2.Score >= InitMaxScore {
			g.Player2.HasWon = true
		
		}

	default:
		// Update Pads
		if strings.Contains(input, "P1") {
			g.Player1.Update(input)
		} else if strings.Contains(input, "P2") {
			g.Player2.Update(input)
		
		}		
	}
}