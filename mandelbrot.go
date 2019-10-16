package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func buildFile(z0 complex128, index uint8) {
	img := buildMandelbrot(z0)
	fileName := fmt.Sprintf("mandelbrot%v.png", index)
	f, _ := os.Create(fileName)
	png.Encode(f, img)
	fmt.Printf("File %v created.\n", fileName)
}

func buildSimpleShots(shots int) {
	for sIndex := 0; sIndex < shots; sIndex++ {
		z0 := complex(0, float64(sIndex)/float64(shots))
		img := buildMandelbrot(z0)
		f, _ := os.Create(fmt.Sprintf("mandelbrot%v.png", sIndex))
		png.Encode(f, img)
		fmt.Printf("File mandelbrot%v.png created.\n", sIndex)
	}
}

func buildMandelbrot(z0 complex128) *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = -1.5, -0.5, -0.5, +0.5
		width, height          = 1024, 1024
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

	return img
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
			//quotient := float64(n) / float64(iterations)
			return color.Gray{255 - contrast*n}
			//return color.RGBA{150 - contrast*n, contrast * n, 255 - contrast*n, 155}
			//return color.RGBA{0, 1, uint8(quotient * float64(255)), 10}

			//size := math.Sqrt(real(v)*real(v) + imag(v)*imag(v))
			//smoothed := math.Log(math.Log(size) * )
		}
	}
	return color.Black
}
