package mcolor

import (
	"image"
	"image/color"
	"math"

	"github.com/dontomato/mandelbrot/pkg/mcalc"
)

const contrast = 15

// CreateGreyRGBA returns grey RGBA with contrast = 15
func CreateGreyRGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelConstParameters) *image.RGBA {
	return createSimpleRGBA(data, settings, func(r mcalc.MandelIterationResult) color.Color {
		if r.Iteration == settings.IterationCount && r.Z == 0 {
			return color.Black
		} else {
			return color.Gray{uint8(settings.IterationCount) - 15*uint8(r.Iteration)}
		}
	})
}

// CreateSmoothSimpleRGBA returns RGBA using primitive smooth technique
func CreateSmoothSimpleRGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelConstParameters) *image.RGBA {
	return createSimpleRGBA(data, settings, func(r mcalc.MandelIterationResult) color.Color {
		if r.Iteration == settings.IterationCount && r.Z == 0 {
			return color.Black
		} else {
			nsmooth := r.Iteration + 1 - int(math.Log(math.Log(r.Z))/math.Log(2))

			return color.RGBA{
				R: 20 + uint8(nsmooth*r.Iteration/2),
				G: 10 + uint8(nsmooth*r.Iteration),
				B: 50 + uint8(nsmooth*r.Iteration),
				A: 255 - uint8(contrast*r.Iteration),
			}
		}
	})
}

func createSimpleRGBA(data []mcalc.MandelIterationResult, settings *mcalc.MandelConstParameters, getColor func(mcalc.MandelIterationResult) color.Color) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, settings.Width, settings.Height))

	i := 0
	for py := 0; py < settings.Height; py++ {
		for px := 0; px < settings.Width; px++ {
			img.Set(px, py, getColor(data[i]))
			i++
		}
	}

	return img
}
