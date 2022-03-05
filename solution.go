package square

// package main

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

type intCustomType int

const (
	SidesTriangle intCustomType = 3
	SidesSquare   intCustomType = 4
	SidesCircle   intCustomType = 0
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	if sidesNum == SidesTriangle {
		return sideLen * sideLen * math.Sqrt(3) / 4
	} else if sidesNum == SidesSquare {
		return sideLen * sideLen
	} else if sidesNum == SidesCircle {
		return math.Pi * sideLen * sideLen
	}

	return 0
}
