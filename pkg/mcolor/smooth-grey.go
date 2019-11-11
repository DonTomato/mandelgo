package mcolor

import (
	"image"
	"image/color"

	"github.com/dontomato/mandelbrot/pkg/mcalc"
)

func getGreyPalette() []color.Color {
	pal := make([]color.Color, 255*3)
	for i := 0; i < 255*3; i++ {
		pal[i] = color.RGBA{uint8(i / 3), uint8((i + 1) / 3), uint8((i + 2) / 3), 255}
	}
	return pal
}

// CreateSmoothGreyRGBA returns smooth RGBA with Grey palette
func CreateSmoothGreyRGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelSettings) *image.RGBA {
	return createPaletteRGBA(data, settings, getGreyPalette())
}
