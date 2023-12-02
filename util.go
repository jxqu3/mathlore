package main

import (
	"image/color"
	"math"
	"os"

	"github.com/d5/tengo/v2"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var tengoScr *tengo.Compiled // We will compile it only once to make the program much faster

// Evaluate the script
func f(x float64) float64 {
	// Set all the variables for the script
	tengoScr.Set("t", t)
	tengoScr.Set("x", x)
	tengoScr.Set("div", div)

	tengoScr.Run()

	return tengoScr.Get("y").Float()
}

func loadScript() []byte {
	content, err := os.ReadFile(Script)
	if err != nil {
		content = []byte("y = x * t") //Default script in case the file isn't open yet
	}
	return []byte("math := import(\"math\")\n" + string(content))
}

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
