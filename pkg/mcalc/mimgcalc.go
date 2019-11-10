package mcalc

import (
	"image"
	"image/color"
	"math/cmplx"
)

// CreateMandelRGBA creates RGBA image of a piece of the Mandelbrot set
func CreateMandelRGBA(params *MandelPictureParameters) *image.RGBA {
	width, height := params.Settings.Width, params.Settings.Height
	realHeight := (float64(height) * params.RealWidth) / float64(width)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*realHeight + (params.Y - params.RealWidth*float64(height)/(2*float64(width)))
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*params.RealWidth + (params.X - params.RealWidth/2)
			z := complex(x, y)
			img.Set(px, py, mandelbrotPoint(z, params))
		}
	}

	return img
}

func mandelbrotPoint(z complex128, params *MandelPictureParameters) color.Color {

	var v = params.Z0
	for n := 0; n < params.Settings.IterationCount; n++ {
		v = v*v + z
		zn := cmplx.Abs(v)
		if zn > 2 {
			return params.GetColorFunc(zn, n)
		}
	}
	return color.Black
}
