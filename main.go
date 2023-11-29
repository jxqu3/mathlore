package main

import (
	"fmt"
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const LineWidth = 1
const Radius = 400

var div = 400  // Number of divisions/lines
var mult = 1.0 // Frame multiplier

var t = 0.0        // The frame counter
var paused = false // Whether the visualizer is paused

var divAngle = 2 * math.Pi / float64(div)

func f(x float64) float64 {
	return x * t
}

func draw() {
	rl.ClearBackground(color.RGBA{20, 20, 20, 255})
	rl.SetLineWidth(LineWidth)

	rl.DrawCircleLines(500, 500, Radius, rl.White)

	for i := 0; i < div; i++ {
		x := float64(i)

		y := f(x)

		startPosX, startPosY := getXY(int(x))
		endPosX, endPosY := getXY(int(y))

		rl.DrawLine(startPosX, startPosY, endPosX, endPosY, rl.White)
	}

	rl.DrawText(fmt.Sprintf("divisions=%d", div), 10, 10, 20, rl.White)
	rl.DrawText(fmt.Sprintf("t=%.2f", t), 10, 40, 20, rl.White)
	rl.DrawText(fmt.Sprintf("multiplier=%.2f", mult), 10, 70, 20, rl.White)
	rl.DrawText(fmt.Sprintf("Paused: %t", paused), 10, 100, 20, rl.White)
	rl.DrawText(fmt.Sprint("FPS: ", rl.GetFPS()), 900, 10, 20, rl.White)
}

func getXY(i int) (int32, int32) {
	x := int32(math.Sin(divAngle*float64(i))*400 + 500)
	y := int32(math.Cos(divAngle*float64(i))*400 + 500)
	return x, y
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
			divAngle = 2 * math.Pi / float64(div)
		}
		if !paused {
			t += float64(rl.GetFrameTime()) * mult
		}

		rl.BeginDrawing()
		draw()
		rl.EndDrawing()
	}
}
