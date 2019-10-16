package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func buildSimpleShots(shots int) {
	for sIndex := 0; sIndex < shots; sIndex++ {
		z0 := complex(0, float64(sIndex)/float64(shots))
		buildMandelbrot(z0, sIndex)
	}
}

func buildMandelbrot(z0 complex128, index int) {
	const (
		xmin, ymin, xmax, ymax = -1.5, -0.5, -0.5, +0.5
		width, height          = 1024 * 2, 1024 * 2
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotPoint(z, z0))
		}
	}

	f, _ := os.Create(fmt.Sprintf("mandelbrot%v.png", index))
	png.Encode(f, img)
}

func mandelbrotPoint(z complex128, z0 complex128) color.Color {
	const (
		iterations = 200
		contrast   = 15
	)

	var v = z0
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
