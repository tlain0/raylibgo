package main

import (
	"fmt"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var screenWidth, screenHeight int32 = 800, 450
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")

	ball := Ball{rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight / 2)}, 10, rl.Vector2{X: randgen(2, 5), Y: randgen(-16, 8)}}

	left := Paddle{rl.Vector2{X: 50, Y: 175}, 10, 100, rl.KeyW, rl.KeyS}
	right := Paddle{rl.Vector2{X: 750, Y: 175}, 10, 100, rl.KeyUp, rl.KeyDown}

	score := []int{0, 0}

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		// Setup
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		// Draw Score
		rl.DrawText(fmt.Sprint(score[0]), 20, 20, 50, rl.Gray)
		rl.DrawText(fmt.Sprint(score[1]), 760, 20, 50, rl.Gray)

		// Draw Sprites
		rl.DrawCircle(int32(ball.Position.X), int32(ball.Position.Y), 10, rl.Black)
		rl.DrawRectangle(int32(left.Position.X), int32(left.Position.Y), left.Thickness, left.Height, rl.Red)
		rl.DrawRectangle(int32(right.Position.X), int32(right.Position.Y), right.Thickness, right.Height, rl.Blue)

		rl.EndDrawing()

		// Check for input every frame
		left.movement()
		right.movement()

		//Move ball and check for collisions
		ball.update()
		ball.wallCollision()
		ball.paddleCollision(left)
		ball.paddleCollision(right)

		// Check for goal
		if ball.Position.X <= 0 {
			score[1] = score[1] + 1
			ball.Position = rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight / 2)}
			ball.Velocity = rl.Vector2{X: randgen(2, 5), Y: randgen(-16, 8)}
		} else if ball.Position.X >= 800 {
			score[0] = score[0] + 1
			ball.Position = rl.Vector2{X: float32(screenWidth / 2), Y: float32(screenHeight / 2)}
			ball.Velocity = rl.Vector2{X: randgen(2, 5), Y: randgen(-16, 8)}
		}

	}

	rl.CloseWindow()
}

type Paddle struct {
	Position  rl.Vector2
	Thickness int32
	Height    int32
	Up        int32
	Down      int32
}

type Ball struct {
	Position rl.Vector2
	Radius   int32
	Velocity rl.Vector2
}

func (ball *Ball) wallCollision() {
	if ball.Position.Y <= 3 {
		ball.Velocity.Y = -ball.Velocity.Y
	}
	if ball.Position.Y >= 450 {
		ball.Velocity.Y = -ball.Velocity.Y
	}
}

func (ball *Ball) paddleCollision(paddle Paddle) {
	var rect rl.Rectangle = rl.Rectangle{X: paddle.Position.X, Y: paddle.Position.Y, Width: float32(paddle.Thickness), Height: float32(paddle.Height)}

	if rl.CheckCollisionCircleRec(ball.Position, float32(ball.Radius), rect) {
		ball.Velocity.X = -ball.Velocity.X
	}
}

func (ball *Ball) update() {
	ball.Position.X += ball.Velocity.X
	ball.Position.Y += ball.Velocity.Y
}

func (paddle *Paddle) movement() {
	// Controls
	if rl.IsKeyDown(paddle.Up) {
		if paddle.Position.Y == 0 {
			paddle.Position.Y = 0
		} else {
			paddle.Position.Y -= 5
		}
	}
	if rl.IsKeyDown(paddle.Down) {
		if paddle.Position.Y >= 350 {
			paddle.Position.Y = 350
		} else {
			paddle.Position.Y += 5
		}
	}
}

func randgen(min float32, max float32) float32 {
	s := rand.NewSource(time.Now().UnixNano())
	x := rand.New(s)
	return (x.Float32() * min) + max
}
