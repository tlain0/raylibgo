package gamelogic

import rl "github.com/gen2brain/raylib-go/raylib"

var cellSize int32 = 25
var gameWidth int32 = 500
var gameHeight int32 = 500

var rowsCount = gameWidth / cellSize
var columnsCount = gameHeight / cellSize

func RenderGrid() {
	for i := int32(0); i < rowsCount; i++ {
		rl.DrawLine(
			int32(0),
			i*cellSize,
			gameWidth,
			i*cellSize,
			rl.Blue,
		)
	}

	for i := int32(0); i < columnsCount; i++ {
		rl.DrawLine(
			i*cellSize,
			int32(0),
			i*cellSize,
			gameHeight,
			rl.Blue,
		)
	}
}
