package mcolor

import (
	"image"
	"image/color"
	"image/color/palette"

	"github.com/dontomato/mandelbrot/pkg/mcalc"
)

// CreateSmoothRetroRGBA returns smooth RGBA with Retro palette
func CreateSmoothRetroRGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelSettings) *image.RGBA {
	return createPaletteRGBA(data, settings, retro)
}

var retro = []color.Color{
	color.RGBA{0x00, 0x04, 0x0f, 0xff},
	color.RGBA{0x03, 0x26, 0x28, 0xff},
	color.RGBA{0x07, 0x3e, 0x1e, 0xff},
	color.RGBA{0x18, 0x55, 0x08, 0xff},
	color.RGBA{0x5f, 0x6e, 0x0f, 0xff},
	color.RGBA{0x84, 0x50, 0x19, 0xff},
	color.RGBA{0x9b, 0x30, 0x22, 0xff},
	color.RGBA{0xb4, 0x92, 0x2f, 0xff},
	color.RGBA{0x94, 0xca, 0x3d, 0xff},
	color.RGBA{0x4f, 0xd5, 0x51, 0xff},
	color.RGBA{0x66, 0xff, 0xb3, 0xff},
	color.RGBA{0x82, 0xc9, 0xe5, 0xff},
	color.RGBA{0x9d, 0xa3, 0xeb, 0xff},
	color.RGBA{0xd7, 0xb5, 0xf3, 0xff},
	color.RGBA{0xfd, 0xd6, 0xf6, 0xff},
	color.RGBA{0xff, 0xf0, 0xf2, 0xff},
}

// CreateSmoothPlan9RGBA returns smooth RGBA with Plan9 palette
func CreateSmoothPlan9RGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelSettings) *image.RGBA {
	return createPaletteRGBA(data, settings, palette.Plan9)
}

// CreateSmoothWebSafeRGBA returns smooth RGBA with WebSafe palette
func CreateSmoothWebSafeRGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelSettings) *image.RGBA {
	return createPaletteRGBA(data, settings, palette.WebSafe)
}
