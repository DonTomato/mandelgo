package mcolor

import (
	"image"

	"github.com/dontomato/mandelbrot/pkg/mcalc"
)

// GetAllFunctions returns map of all color functions
func GetAllFunctions() map[string]func(result []mcalc.MandelIterationResult, settings *mcalc.MandelConstParameters) *image.RGBA {
	return map[string]func(result []mcalc.MandelIterationResult, settings *mcalc.MandelConstParameters) *image.RGBA{
		"grey":   CreateGreyRGBA,
		"smooth": CreateSmoothSimpleRGBA,
	}
}
