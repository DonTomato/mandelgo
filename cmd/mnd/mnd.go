package main

import (
	"fmt"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/dontomato/mandelbrot/pkg/mcalc"
	"github.com/dontomato/mandelbrot/pkg/mconfig"
)

func main() {
	fmt.Println("Mandelbrot set experiments")
	conf, err := mconfig.Get()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Conf: %v\n", conf.DataPath)

	settings := mcalc.MandelConstParameters{Width: 2560, Height: 1440, IterationCount: 200}

	params := mcalc.MandelPictureParameters{
		GetColorFunc: func(z float64, iteration int) color.Color {
			return color.Gray{255 - 15*uint8(iteration)}
		},
		RealWidth: 4,
		Settings:  &settings,
		X:         0,
		Y:         0,
		Z0:        complex(0, 0),
	}

	img := mcalc.CreateMandelRGBA(&params)

	fileName := filepath.Join(conf.DataPath, fmt.Sprintf("mandelbrot%v.jpg", 0))
	f, err := os.Create(fileName)
	if err != nil {
		if conf.LogEnable {
			log.Fatal(err)
		}
		return
	}

	var opt jpeg.Options
	opt.Quality = 80
	jpeg.Encode(f, img, &opt)

	fmt.Printf("File %v created.\n", fileName)
}
