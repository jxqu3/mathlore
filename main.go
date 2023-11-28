package main

import (
	"fmt"
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var t = 0.0        // The frame counter
var mult = 1.0     // Frame multiplier
var div = 20       // Number of divisions/lines
var paused = false // Whether the visualizer is paused

func draw() {
	rl.ClearBackground(color.RGBA{20, 20, 20, 255})
	rl.SetLineWidth(2)

	rl.DrawCircleLines(500, 500, 400, rl.White)

	for i := 0; i < div; i++ {
		res := float64(i) / float64(div) * t
		x := int32(math.Sin(res)*400 + 500)
		y := int32(math.Cos(res)*400 + 500)
		rl.DrawLine(500, 500, x, y, rl.White)
	}
	rl.DrawText(fmt.Sprintf("divisions=%d", div), 10, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("t=%.2f", t), 10, 40, 20, rl.White)
	rl.DrawText(fmt.Sprintf("multiplier=%.2f", mult), 10, 70, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Paused: %t", paused), 10, 100, 20, rl.White)
	rl.DrawText(fmt.Sprint("FPS: ", rl.GetFPS()), 900, 10, 20, rl.White)
}

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint | rl.FlagWindowHighdpi | rl.FlagVsyncHint)
	rl.InitWindow(1000, 1000, "circlesim")

	for !rl.WindowShouldClose() {
		switch {
		case rl.IsKeyPressed(rl.KeySpace):
			paused = !paused
		case rl.IsKeyDown(rl.KeyLeftShift):
			mult += float64(rl.GetMouseWheelMove()) * 0.1
		default:
			div += int(rl.GetMouseWheelMove())
		}
		if !paused {
			t += float64(rl.GetFrameTime()) * mult
		}
		rl.BeginDrawing()
		draw()
		rl.EndDrawing()
	}
}
