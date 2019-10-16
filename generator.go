package main

import (
	"runtime"
)

type initialCondition struct {
	z0    complex128
	index uint8
}

func generate() {
	const number = 100

	queue := newQueue(number)
	for i := 0; i < number; i++ {
		queue.push(&initialCondition{z0: complex(0, float64(i)/float64(number)), index: uint8(i)})
	}
	//fmt.Printf("Slice: %v\n", queue.Items)
	for thIndex := 0; thIndex < runtime.GOMAXPROCS(0); thIndex++ {
		go process(queue)
	}
}

func process(queue *concurrentQueue) {
	for item := queue.pop(); item != nil; item = queue.pop() {
		//fmt.Printf("Item: %v;\n", item)
		buildFile(item.z0, item.index)
	}
}
