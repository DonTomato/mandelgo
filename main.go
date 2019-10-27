package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	//xmin, ymin, xmax, ymax = 0.0, -0.7, 0.5, -0.4
	xCenter, yCenter = -0.15, -0.68
	xSize            = 0.8

	width, height = 2560, 1440

	r0, r1 = 0.5, 1.0
)

const filesCount = 2000

//const width = 3072

func main() {
	fmt.Println("Mandelbrot set experiments")
	reader := bufio.NewReader(os.Stdin)

	start := time.Now()

	//fmt.Println(os.Args[0])
	//dir, _ := os.Getwd()
	//fmt.Println(dir)

	//if len(os.Args) > 1 && os.Args[1] =
	generate()
	generateAvi()

	fmt.Printf("Process finished. Time: %v\n", time.Now().Sub(start))

	reader.ReadString('\n')
}
