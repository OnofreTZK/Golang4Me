package main

import (
	"math"
)

// Initializing with Capitalized letter is PUBLIC(visible in and out of the package)
// Initializing with lower letter is PRIVATE(visible only in the package)

// Example:
// Name -> Public
// name or _name -> Private

// Coordinates of a straight in the cartesian plane
type Point struct {
	x float64
	y float64
}

func catetos(a, b Point) (cx, cy float64) {

	cx = math.Abs(b.x - a.x)
	cy = math.Abs(b.y - a.y)

	return // cx cy
}

// Linear distance between 2 points
func Distance(a, b Point) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}
