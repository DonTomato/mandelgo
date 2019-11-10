package main

import (
	"fmt"

	"github.com/dontomato/mandelbrot/pkg/mconfig"
)

func main() {
	//conf := mcon
	conf := mconfig.Get()
	fmt.Println("Mandelbrot set experiments")
	fmt.Printf("Conf: %v\n", conf.DataPath)
}
