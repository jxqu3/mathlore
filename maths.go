package main

import "math"

// The function to be visualized
// f(x) =
func f(x float64) float64 {
	return math.Sin(t) * x
}

// Constants
const LineWidth = 1
const FontName = "res/playfair.ttf"
const Radius = 400

// Global variables
var div = 400  // Number of divisions/lines
var mult = 0.1 // Frame multiplier
