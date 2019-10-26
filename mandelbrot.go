package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"math/cmplx"
	"os"
	"path/filepath"
)

func buildFile(z0 complex128, index uint8) {
	img := buildMandelbrot(z0)
	fileName := filepath.Join("data", fmt.Sprintf("mandelbrot%v.jpg", index))
	f, _ := os.Create(fileName)
	//png.Encode(f, img)

	var opt jpeg.Options
	opt.Quality = 80
	jpeg.Encode(f, img, &opt)

	fmt.Printf("File %v created.\n", fileName)
}

func buildSimpleShots(shots int) {
	for sIndex := 0; sIndex < shots; sIndex++ {
		z0 := complex(0, float64(sIndex)/float64(shots))
		img := buildMandelbrot(z0)
		f, _ := os.Create(filepath.Join("data", fmt.Sprintf("mandelbrot%v.jpg", sIndex)))
		png.Encode(f, img)
		fmt.Printf("File mandelbrot%v.png created.\n", sIndex)
	}
}

func buildMandelbrot(z0 complex128) *image.RGBA {
	const (
		xmin, ymin, xmax, ymax = 0.0, -0.7, 0.5, -0.4
		//xmin, ymin, xmax, ymax = -2, -1.5, 1, 1.5
		//width, height          = 1024, 1024
	)

	width := 2048
	height := int(math.Trunc((ymax - ymin) * float64(2048) / (xmax - xmin)))

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrotPoint(z, z0))
		}
	}

	return img
}

func mandelbrotPoint(z complex128, z0 complex128) color.Color {
	const (
		iterations = 255
		contrast   = 5
	)

	var v = z0
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		zn := cmplx.Abs(v)
		if zn > 2 {

			//return color.Gray{255 - contrast*n}

			nsmooth := n + 1 - uint8(math.Log(math.Log(zn))/math.Log(2))
			return color.RGBA{150 + nsmooth, 0, 54 + nsmooth, 230 - contrast*n}

			// r2 := math.Pow(cmplx.Abs(v), 2)

			// if r2 > 1000000 {
			// 	vk := math.Log(r2)

			// 	rgbV := 255 * (1 + math.Cos(2*math.Pi*vk)) / 2
			// 	rz := uint8(rgbV)
			// 	return color.RGBA{rz, rz, rz, 0}
			// }
		}
	}
	return color.Black
}
