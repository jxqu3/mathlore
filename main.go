package main

import (
	"fmt"
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var frame = 0.0
var mult = 0.1
var div = 200
var paused = false

func draw() {
	rl.ClearBackground(color.RGBA{20, 20, 20, 255})
	rl.SetLineWidth(2)

	rl.DrawCircleLines(500, 500, 400, rl.White)

	for i := 0; i < div; i++ {
		res := frame * float64(i)
		x := int32(math.Cos(res)*400 + 500)
		y := int32(math.Sin(res)*400 + 500)
		rl.DrawLine(218, 218, x, y, rl.White)
	}
	rl.DrawText(fmt.Sprintf("x=%.2f", frame), 10, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("mult=%.2f", mult), 10, 40, 20, rl.White)
	rl.DrawText(fmt.Sprint("FPS: ", rl.GetFPS()), 10, 70, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Paused: %t", paused), 10, 100, 20, rl.White)
}

func main() {
	rl.SetConfigFlags(rl.FlagMsaa4xHint | rl.FlagWindowHighdpi | rl.FlagVsyncHint)
	rl.InitWindow(1000, 1000, "Ventanica")

	for !rl.WindowShouldClose() {
		if rl.IsKeyPressed(rl.KeySpace) {
			paused = !paused
		}
		if !paused {
			frame += float64(rl.GetFrameTime()) * mult
			mult += float64(rl.GetMouseWheelMove()) * 0.1
		}
		rl.BeginDrawing()
		draw()
		rl.EndDrawing()
	}
}
