package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"

	"github.com/dontomato/mandelbrot/pkg/mcalc"
	"github.com/dontomato/mandelbrot/pkg/mcolor"
	"github.com/dontomato/mandelbrot/pkg/mconfig"
)

func main() {
	fmt.Println("Mandelbrot set experiments")

	functions := mcolor.GetAllFunctions()
	colorFunc := "grey"

	if len(os.Args) > 1 {
		if _, ok := functions[os.Args[1]]; ok {
			colorFunc = os.Args[1]
		} else {
			log.Printf("Error: %v doesn't exist; Grey func was used instead.", os.Args[1])
		}
	}

	conf, err := mconfig.Get()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
	fmt.Printf("Conf: %v\n", conf.DataPath)

	if _, err := os.Stat(conf.DataPath); os.IsNotExist(err) {
		log.Fatal("Data path doesn't exist")
		return
	}

	settings := mcalc.MandelConstParameters{Width: 2560, Height: 1440, IterationCount: 200}

	params := mcalc.MandelPictureParameters{
		//GetColorFunc: functions[colorFunc],
		RealWidth:  4,
		Settings:   &settings,
		X:          -0.7,
		Y:          0,
		Z0:         complex(0, 0),
		CreateRGBA: functions[colorFunc],
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
