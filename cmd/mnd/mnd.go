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

	settings := mcalc.MandelSettings{Width: 2560, Height: 1440, IterationCount: 200}

	params1 := mcalc.MandelPictureParameters{
		RealWidth:  4,
		Settings:   &settings,
		X:          0,
		Y:          0,
		Z0:         complex(0, 0),
		CreateRGBA: functions[colorFunc],
	}

	params2 := mcalc.MandelPictureParameters{
		RealWidth:  0.3,
		Settings:   &settings,
		X:          -1.25,
		Y:          -0.1,
		Z0:         complex(0, 0),
		CreateRGBA: functions[colorFunc],
	}

	createFile(&params1, conf.DataPath, 0)
	createFile(&params2, conf.DataPath, 1)
}

func createFile(params *mcalc.MandelPictureParameters, path string, index int) {
	img := mcalc.CreateMandelRGBA(params)

	fileName := filepath.Join(path, fmt.Sprintf("mandelbrot%v.jpg", index))
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}

	var opt jpeg.Options
	opt.Quality = 80
	jpeg.Encode(f, img, &opt)

	fmt.Printf("File %v created.\n", fileName)
}
