package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Mandelbrot set experiments")
	reader := bufio.NewReader(os.Stdin)

	buildSimpleShots(10)

	fmt.Println("Process finished.")

	reader.ReadString('\n')
}
