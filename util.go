package main

import (
	"image/color"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func drawCircle(cx, cy, radius float32, step float64, color color.RGBA) {
	rl.Begin(rl.Lines)
	for angle := 0.0; angle < 360; angle += step {
		rl.Vertex2f(cx+float32(math.Cos(rl.Deg2rad*angle))*radius, cy+float32(math.Sin(rl.Deg2rad*angle))*radius)
		rl.Vertex2f(cx+float32(math.Cos(rl.Deg2rad*(angle+step)))*radius, cy+float32(math.Sin(rl.Deg2rad*(angle+step)))*radius)
	}
}

const FontSize = 26

func drawText(text string, x, y int32, color color.RGBA) {
	rl.DrawTextEx(font, text, rl.Vector2{X: float32(x), Y: float32(y)}, float32(FontSize), 0, color)
}
