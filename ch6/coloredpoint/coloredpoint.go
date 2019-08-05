package coloredpoint

import "image/color"

//结构体组合

type Point struct {
	X,Y float64
}

type ColoredPoint struct {
	Point
	Color color.RGBA
}
