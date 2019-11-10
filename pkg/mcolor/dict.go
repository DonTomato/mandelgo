package mcolor

import (
	"image/color"
)

// GetAllFunctions returns map of all color functions
func GetAllFunctions() map[string]func(z float64, iteration int) color.Color {
	return map[string]func(z float64, iteration int) color.Color{
		"grey":   GetGrayPoint,
		"smooth": GetSmoothForPoint,
	}
}
