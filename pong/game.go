package pong

type Game struct {
	*Ball
	Player1  *Paddle
	Player2  *Paddle
	MaxScore int
}

const (
	initBallVelocity = 5.0
	initPaddleSpeed  = 10.0
	speedUpdateCount = 6
	speedIncrement   = 0.5

	ScreenWidth = 600
	ScreenHeight = 600
)

func (g *Game) init() {
	g.MaxScore = 10

	g.Player1 = &Paddle{
		Position: Position{
			X: 10,
			Y: 10,
		},
		Score: 0,
		Speed: initPaddleSpeed,
		Width: InitPaddleWidth,
		Height: InitPaddleHeight,
	}

	g.Player2 = &Paddle{
		Position: Position{
			X: 20,
			Y: 20,
		},
		Score: 0,
		Speed: initPaddleSpeed,
		Width: InitPaddleWidth,
		Height: InitPaddleHeight,
	}

	g.Ball = &Ball{
		Position: Position{
			X: 30,
			Y: 30,
		},

		Radius: InitBallRadius,
		XVelocity: initBallVelocity,
		YVelocity: initBallVelocity,
	}
}