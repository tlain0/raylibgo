package gui

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Button struct {
	rl.Rectangle
}

func NewButton(x float32, y float32, width float32, height float32, text string) bool {
	var button = Button{rl.Rectangle{x, y, width, height}}

	// Paint the button at XY coordinates
	rl.DrawRectangle(int32(button.Rectangle.X), int32(button.Rectangle.Y), int32(button.Rectangle.Width), int32(button.Rectangle.Height), rl.Green)

	// Draw text at XY coordinates
	//rl.DrawText(text, (int32-rl.MeasureText(text, 30))/2, 400, 30, rl.Black)
	rl.DrawText(text, int32(x), int32(y), 30, rl.Black)

	// Check for clicks
	mousePoint := rl.GetMousePosition()
	if (rl.CheckCollisionPointRec(mousePoint, button.Rectangle)) && rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		fmt.Println(text)
		return true
	}
	return false

}
