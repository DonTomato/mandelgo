package main

type initialCondition struct {
	z0    complex128
	index uint8
}

func generate() {
	const number = 50

	queue := NewQueue(number)
	for i := 0; i < number; i++ {
		queue.Items[i] = initialCondition{z0: complex(0, float64(i)/float64(number)), index: uint8(i)}
	}

}
