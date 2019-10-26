package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Mandelbrot set experiments")
	reader := bufio.NewReader(os.Stdin)

	start := time.Now()

	fmt.Println(os.Args[0])

	generate()
	generateAvi()

	fmt.Printf("Process finished. Time: %v\n", time.Now().Sub(start))

	reader.ReadString('\n')
}
