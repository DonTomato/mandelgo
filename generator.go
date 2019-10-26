package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/icza/mjpeg"
)

type initialCondition struct {
	z0    complex128
	index uint16
}

func generate() {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		os.Mkdir("data", os.ModePerm)
	}

	queue := newQueue(filesCount)
	for i := 0; i < filesCount; i++ {
		queue.push(&initialCondition{z0: complex(0, float64(i)/float64(filesCount)), index: uint16(i)})
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
	aw, err := mjpeg.New(filepath.Join("data", "mandelbrot.avi"), width, int32(getHeigh()), 20)
	checkErr(err)

	for i := 0; i < filesCount; i++ {
		data, err := ioutil.ReadFile(filepath.Join("data", fmt.Sprintf("./mandelbrot%v.jpg", i)))
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
