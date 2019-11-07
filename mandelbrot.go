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

	xSizeV := xSize - xSize*(0.90)*float64(index)/float64(filesCount)

	img := buildMandelbrot(z0, xSizeV)
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
		img := buildMandelbrot(z0, xSize)
		f, _ := os.Create(filepath.Join("data", fmt.Sprintf("mandelbrot%v.jpg", sIndex)))
		png.Encode(f, img)
		fmt.Printf("File mandelbrot%v.png created.\n", sIndex)
	}
}

func buildMandelbrot(z0 complex128, xSizeV float64) *image.RGBA {
	ySize := (height * xSizeV) / width

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*ySize + (yCenter - xSizeV*float64(height)/(2*float64(width)))
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*xSizeV + (xCenter - xSizeV/2)
			z := complex(x, y)
			img.Set(px, py, mandelbrotPoint(z, z0))
		}
	}

	return img
}

const (
	iterations = 200
	contrast   = 15
)

func mandelbrotPoint(z complex128, z0 complex128) color.Color {

	var v = z0
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		zn := cmplx.Abs(v)
		if zn > 2 {

			//return getGrayPoint(zn, n)

			return getColorForPoint(zn, n)
		}
	}
	return color.Black
}

func getGrayPoint(zn float64, n uint8) color.Color {
	return color.Gray{255 - contrast*n}
}

func getColorForPoint(zn float64, n uint8) color.Color {
	nsmooth := n + 1 - uint8(math.Log(math.Log(zn))/math.Log(2))
	return color.RGBA{20 + nsmooth*n/2, 10 + nsmooth*n, 0 + nsmooth*n, 255 - contrast*n}
}
