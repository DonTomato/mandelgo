package mcolor

import (
	"image"
	"image/color"

	"github.com/dontomato/mandelbrot/pkg/mcalc"
)

func getBWPalette() []color.Color {
	blackWhite := make([]color.Color, 0)
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			blackWhite = append(blackWhite, color.RGBA{0, 0, 0, 255})
		} else {
			blackWhite = append(blackWhite, color.RGBA{255, 255, 255, 255})
		}
	}
	return blackWhite
}

// CreateSmoothBlackWhiteRGBA returns smooth RGBA with Grey palette
func CreateSmoothBlackWhiteRGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelSettings) *image.RGBA {
	return createPaletteRGBA(data, settings, getBWPalette())
}
