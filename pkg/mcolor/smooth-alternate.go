package mcolor

import (
	"image"
	"image/color"

	"github.com/dontomato/mandelbrot/pkg/mcalc"
)

func getAlternatePalette() []color.Color {
	alternate := make([]color.Color, 20)
	for i := 0; i < len(alternate); i++ {
		switch i % 6 {
		case 0:
			alternate[i] = color.RGBA{0x18, 0x4d, 0x68, 255}
		case 1:
			alternate[i] = color.RGBA{0x31, 0x80, 0x9f, 255}
		case 2:
			alternate[i] = color.RGBA{0xfb, 0x9c, 0x6c, 255}
		case 3:
			alternate[i] = color.RGBA{0xd5, 0x51, 0x21, 255}
		case 4:
			alternate[i] = color.RGBA{0xcf, 0xe9, 0x90, 255}
		case 5:
			alternate[i] = color.RGBA{0xea, 0xfb, 0xc5, 255}
		}
	}
	return alternate
}

// CreateSmoothAlternateRGBA returns smooth RGBA with Grey palette
func CreateSmoothAlternateRGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelSettings) *image.RGBA {
	return createPaletteRGBA(data, settings, getAlternatePalette())
}
