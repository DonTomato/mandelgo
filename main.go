package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	xmin, ymin, xmax, ymax = 0.0, -0.7, 0.5, -0.4
	//xmin, ymin, xmax, ymax = -1.5, -0.8, 0, 0.3
)

const filesCount = 1000
const width = 3072

func main() {
	fmt.Println("Mandelbrot set experiments")
	reader := bufio.NewReader(os.Stdin)

	start := time.Now()

	fmt.Println(os.Args[0])
	dir, _ := os.Getwd()
	fmt.Println(dir)

	generate()
	generateAvi()

	fmt.Printf("Process finished. Time: %v\n", time.Now().Sub(start))

	reader.ReadString('\n')
}
