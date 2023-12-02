package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/d5/tengo/v2"
	"github.com/d5/tengo/v2/stdlib"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func renderHUD() {
	if showHUD {
		drawText(fmt.Sprintf("divisions=%d", div), 10, 10, rl.White)
		drawText(fmt.Sprintf("t=%.2f", t), 10, 40, rl.White)
		drawText(fmt.Sprintf("multiplier=%.2f", mult), 10, 70, rl.White)
		drawText(fmt.Sprintf("f(25)=%.2f", f(25)), 10, 100, rl.White)
		drawText(fmt.Sprintf("Paused: %t", paused), 10, 130, rl.White)

		drawText(fmt.Sprint("FPS: ", rl.GetFPS()), 900, 10, rl.White)
	}
}

func draw() {
	rl.ClearBackground(color.RGBA{20, 20, 20, 255})
	rl.SetLineWidth(LineWidth)

	drawCircle(500, 500, Radius, 3, rl.White)

	for i := 0; i < div; i++ {
		x := float64(i)

		y := f(x)

		startPosX, startPosY := getXY(x)
		endPosX, endPosY := getXY(y)

		rl.DrawLine(startPosX, startPosY, endPosX, endPosY, rl.White)
	}

}

func handleInput(cam *rl.Camera2D) {
	switch {
	case rl.IsKeyPressed(rl.KeySpace):
		paused = !paused
	case rl.IsKeyPressed(rl.KeyF1):
		showHUD = !showHUD
	case rl.IsKeyDown(rl.KeyLeftShift):
		if rl.IsKeyDown(rl.KeyLeftAlt) {
			mult += float64(rl.GetMouseWheelMove()) * 2
		} else {
			mult += float64(rl.GetMouseWheelMove()) * 0.1
		}
	case rl.IsKeyDown(rl.KeyLeftControl):
		cam.Zoom = max(cam.Zoom+float32(rl.GetMouseWheelMove()*0.1), 0.1)
	default:
		div += int(rl.GetMouseWheelMove())
	}
	if !paused {
		t += float64(rl.GetFrameTime()) * mult
	}

	mouseDelta := rl.GetMouseDelta()

	if rl.IsMouseButtonDown(rl.MouseButtonMiddle) {
		cam.Offset.X += mouseDelta.X
		cam.Offset.Y += mouseDelta.Y
	}

}

func getXY(i float64) (int32, int32) {
	divAngle := 2 * math.Pi / float64(div)

	x := int32(math.Sin(divAngle*i)*400 + 500)
	y := int32(math.Cos(divAngle*i)*400 + 500)

	return x, y
}
func main() {
	go func() {
		scr := tengo.NewScript(loadScript())
		scr.SetImports(stdlib.GetModuleMap("math"))
		scr.Add("x", 0)
		scr.Add("t", 0)
		scr.Add("div", div)
		compiled, _ := scr.Compile()
		tengoScr = compiled
	}()

	rl.SetConfigFlags(rl.FlagMsaa4xHint | rl.FlagVsyncHint | rl.FlagWindowHighdpi)
	rl.InitWindow(1000, 1000, "circlesim")

	font = rl.LoadFontEx(FontName, FontSize, nil)

	cam := rl.Camera2D{
		Target:   rl.NewVector2(0, 0),
		Zoom:     1.0,
		Offset:   rl.NewVector2(0, 0),
		Rotation: 0,
	}

	for !rl.WindowShouldClose() {
		handleInput(&cam)

		rl.BeginDrawing()
		rl.BeginMode2D(cam)
		draw()
		rl.EndMode2D()
		renderHUD()
		rl.EndDrawing()
	}

	rl.UnloadFont(font)
}
