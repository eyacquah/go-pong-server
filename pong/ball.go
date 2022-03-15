package pong

type Ball struct {
	Position
	Radius    float32
	XVelocity float32
	YVelocity float32
}

const (
	InitBallRadius = 10
)

func (b *Ball) Update(leftPaddle *Paddle, rightPaddle *Paddle) {

	var h = ScreenHeight

	b.X += b.XVelocity
	b.Y += b.YVelocity

	// bounce off edges when getting to top/bottom
	if b.Y-b.Radius > float32(h) {
		b.YVelocity = -b.YVelocity
		b.Y = float32(h) - b.Radius
	} else if b.Y+b.Radius < 0 {
		b.YVelocity = -b.YVelocity
		b.Y = b.Radius
	}

	// bounce off paddles
	if b.X-b.Radius < leftPaddle.X+float32(leftPaddle.Width/2) &&
		b.Y > leftPaddle.Y-float32(leftPaddle.Height/2) &&
		b.Y < leftPaddle.Y+float32(leftPaddle.Height/2) {
		b.XVelocity = -b.XVelocity
		b.X = leftPaddle.X + float32(leftPaddle.Width/2) + b.Radius
	} else if b.X+b.Radius > rightPaddle.X-float32(rightPaddle.Width/2) &&
		b.Y > rightPaddle.Y-float32(rightPaddle.Height/2) &&
		b.Y < rightPaddle.Y+float32(rightPaddle.Height/2) {
		b.XVelocity = -b.XVelocity
		b.X = rightPaddle.X - float32(rightPaddle.Width/2) - b.Radius
	}
}