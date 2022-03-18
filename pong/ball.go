package pong

type Ball struct {
	Position
	Radius    int `json:"radius"`
	XVelocity int `json:"xVelocity"`
	YVelocity int `json:"yVelocity"`
}

const (
	InitBallRadius = 10
)

func (b *Ball) Update(leftPaddle *Paddle, rightPaddle *Paddle) {

	h := ScreenHeight

	b.X += b.XVelocity
	b.Y += b.YVelocity

	// bounce off edges when getting to top/bottom
	if b.Y-b.Radius > h {
		b.YVelocity = -b.YVelocity
		b.Y = h - b.Radius

	} else if b.Y+b.Radius < 0 {

		b.YVelocity = -b.YVelocity
		b.Y = b.Radius
	}

	// Bounce off left wall
	if b.X+b.Radius < 0 {
		b.Position.X = 0
		b.XVelocity *= -1
	}

	// Bounce off right wall
	if b.X > ScreenWidth {
		b.Position.X = ScreenWidth
		b.XVelocity *= -1
	}

	// Bounce off left Paddle

	if b.X-b.Radius < leftPaddle.X+leftPaddle.Width/2 && b.Y > leftPaddle.Y-leftPaddle.Height/2 && b.Y < leftPaddle.Y+leftPaddle.Height {
		b.XVelocity = -b.XVelocity
		b.X = leftPaddle.X + leftPaddle.Width/2 + b.Radius

		if b.X+b.Radius > 0 {
			leftPaddle.Score += 1
		}
	}

	// Bounce off right paddle

	if (b.X+b.Radius > rightPaddle.X-rightPaddle.Width/2) && b.Y > rightPaddle.Y-rightPaddle.Height/2 && b.Y < rightPaddle.Y+rightPaddle.Height {
		b.XVelocity = -b.XVelocity
		b.X = rightPaddle.X - rightPaddle.Width/2 - b.Radius

		// Ensure ball is not beyond the pads before incrementing score
		if b.X < ScreenWidth {
			rightPaddle.Score += 1
		}
	}
}