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

func buildFile(z0 complex128, index uint16) {
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

// func getHeigh() int {
// 	return int(math.Trunc((ymax - ymin) * float64(width) / (xmax - xmin)))
// }

func buildMandelbrot(z0 complex128) *image.RGBA {
	//height := getHeigh()

	ySize := (height * xSize) / width

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*ySize + yLeft
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*xSize + xLeft
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
		zn := cmplx.Abs(v)
		if zn > 2 {

			//return color.Gray{255 - contrast*n}

			nsmooth := n + 1 - uint8(math.Log(math.Log(zn))/math.Log(2))
			return color.RGBA{20 + nsmooth*n/2, 10 + nsmooth*n, 0 + nsmooth*n, 255 - contrast*n}

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
