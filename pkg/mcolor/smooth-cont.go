package mcolor

import (
	"image"
	"image/color"

	"github.com/dontomato/mandelbrot/pkg/mcalc"
)

func getContPalette() []color.Color {
	cont := make([]color.Color, 10000)
	for i := range cont {
		//val := float64(i) / float64(len(cont))
		val := i * 256 / len(cont)
		cont[i] = color.RGBA{uint8(val), 0, uint8(255 - val), uint8(255)}
	}
	return cont
}

// CreateSmoothContRGBA returns smooth RGBA with Grey palette
func CreateSmoothContRGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelSettings) *image.RGBA {
	return createPaletteRGBA(data, settings, getContPalette())
}
