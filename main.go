package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	//xmin, ymin, xmax, ymax = 0.0, -0.7, 0.5, -0.4
	xLeft, yLeft = -1, -0.3
	xSize        = 0.33

	width, height = 2560, 1140
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
