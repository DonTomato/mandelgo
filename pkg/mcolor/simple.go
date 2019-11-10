package mcolor

import (
	"image/color"
	"math"
)

const contrast = 15

// GetGrayPoint - returns grey color point
func GetGrayPoint(z float64, iteration int) color.Color {
	return color.Gray{255 - contrast*uint8(iteration)}
}

// GetSmoothForPoint returns smoosh value
func GetSmoothForPoint(z float64, iteration int) color.Color {
	nsmooth := iteration + 1 - int(math.Log(math.Log(z))/math.Log(2))

	return color.RGBA{
		R: 20 + uint8(nsmooth*iteration/2),
		G: 10 + uint8(nsmooth*iteration),
		B: 50 + uint8(nsmooth*iteration),
		A: 255 - uint8(contrast*iteration),
	}
}
