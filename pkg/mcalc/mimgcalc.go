package mcalc

import (
	"image"
	"math/cmplx"
)

// CreateMandelRGBA creates RGBA image of a piece of the Mandelbrot set
func CreateMandelRGBA(params *MandelPictureParameters) *image.RGBA {
	width, height := params.Settings.Width, params.Settings.Height
	realHeight := (float64(height) * params.RealWidth) / float64(width)

	//img := image.NewRGBA(image.Rect(0, 0, width, height))

	result := make([]MandelIterationResult, width*height)

	i := 0
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*realHeight + (params.Y - params.RealWidth*float64(height)/(2*float64(width)))
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*params.RealWidth + (params.X - params.RealWidth/2)
			z := complex(x, y)
			//img.Set(px, py, mandelbrotPoint(z, params))
			result[i] = mandelbrotPoint(z, params)
			i++
		}
	}

	return params.CreateRGBA(result, params.Settings)
}

func mandelbrotPoint(z complex128, params *MandelPictureParameters) MandelIterationResult {

	var v = params.Z0
	for n := 0; n < params.Settings.IterationCount; n++ {
		v = v*v + z
		zn := cmplx.Abs(v)
		if zn > 2 {
			return MandelIterationResult{Iteration: n, Z: zn}
		}
	}
	return MandelIterationResult{Iteration: params.Settings.IterationCount, Z: 0}
}
