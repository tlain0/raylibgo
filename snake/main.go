package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	gl "raylibcom.com/main/snake/gameLogic"
	"raylibcom.com/main/snake/gui"
)

// Control Game state
type GameState int

const (
	GAME_OVER    GameState = 0
	GAME_RUNNING GameState = 1
	GAME_MENU    GameState = 2
)

var state = GAME_MENU

func main() {
	var screenWidth = 900
	var screenHeight = 720
	var center rl.Vector2 = rl.Vector2{float32(screenWidth / 2), float32(screenHeight / 2)}

	// INIT
	rl.InitWindow(int32(screenWidth), int32(screenHeight), "raylib [core] example - basic window")

	// Load textures before beginning drawing
	var titleScreenText rl.Texture2D = rl.LoadTexture("resources/SnakeTitle.png")
	var apple rl.Texture2D = rl.LoadTexture("resources/Apple.png")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		switch state {
		case 0:
		case 1:
			gamePaintLoop()
		case 2:
			menuPaintLoop(titleScreenText, center, apple)
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}

func menuPaintLoop(titleScreenText rl.Texture2D, center rl.Vector2, apple rl.Texture2D) {
	rl.ClearBackground(rl.RayWhite)
	rl.DrawTexture(titleScreenText, 0, 0, rl.White)
	rl.DrawTexture(apple, 0, 500, rl.White)

	if gui.NewButton((center.X - 150), center.Y, 250, 50, "START GAME") {
		state = GAME_RUNNING
	}
	gui.NewButton((center.X - 150), center.Y+100, 250, 50, "OPTIONS")
}

func gamePaintLoop() {
	rl.ClearBackground(rl.RayWhite)
	gl.RenderGrid()
}
