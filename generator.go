package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/icza/mjpeg"
)

type initialCondition struct {
	z0    complex128
	index uint8
}

const number = 250

func generate() {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", os.ModePerm)
	}

	queue := newQueue(number)
	for i := 0; i < number; i++ {
		queue.push(&initialCondition{z0: complex(0, float64(i)/float64(number)), index: uint8(i)})
	}
	done := make(chan uint8)
	//fmt.Printf("Slice: %v\n", queue.Items)
	for thIndex := 0; thIndex < runtime.GOMAXPROCS(0); thIndex++ {
		go func() {
			for item := queue.pop(); item != nil; item = queue.pop() {
				//fmt.Printf("Item: %v;\n", item)
				buildFile(item.z0, item.index)
			}
			done <- 1
		}()
	}

	for thIndex := 0; thIndex < runtime.GOMAXPROCS(0); thIndex++ {
		<-done
	}
}

func generateAvi() {
	aw, err := mjpeg.New("mandelbrot.avi", 2048, 1228, 60)
	checkErr(err)

	for i := 0; i < number; i++ {
		data, err := ioutil.ReadFile(fmt.Sprintf("./mandelbrot%v.jpg", i))
		checkErr(err)
		aw.AddFrame(data)
	}
}

func process(queue *concurrentQueue) {
	for item := queue.pop(); item != nil; item = queue.pop() {
		//fmt.Printf("Item: %v;\n", item)
		buildFile(item.z0, item.index)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
