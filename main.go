package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	//xmin, ymin, xmax, ymax = 0.0, -0.7, 0.5, -0.4
	xLeft, yLeft = -0.78, -0.3
	xSize        = 0.11

	width, height = 2560, 1140

	r0, r1 = 0.46, 0.72
)

const filesCount = 1000

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
