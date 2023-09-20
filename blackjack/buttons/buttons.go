package logic

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	rl.Rectangle
}

var HitButton = Button{rl.Rectangle{570, 200, 70, 50}}
var StayButton = Button{rl.Rectangle{650, 200, 70, 50}}
var RestartButton = Button{rl.Rectangle{410, 200, 150, 50}}
var GameOverRec = Button{rl.Rectangle{0, 400, 800, 50}}

func (button *Button) DrawButton(text string) {
	rl.DrawRectangle(int32(button.Rectangle.X), int32(button.Rectangle.Y), int32(button.Rectangle.Width), int32(button.Rectangle.Height), rl.Green)
	rl.DrawText(text, int32(button.Rectangle.X), int32(button.Rectangle.Y), 30, rl.Black)
}

func (button *Button) CheckClick(text string) bool {
	mousePoint := rl.GetMousePosition()
	if (rl.CheckCollisionPointRec(mousePoint, button.Rectangle)) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		fmt.Println(text)
		return true
	}
	return false
}
