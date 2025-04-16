package main

import (
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)
	wp := &WorkerPool{
		NumberOfWorkers: 5,
		Tasks:           createTasks(20),
	}
	wp.Run()
}

func createTasks(numOfTasks int) []TaskProcessor {
	tp := make([]TaskProcessor, 0)

	for i := 0; i < numOfTasks; i++ {
		task := &Image{
			Id: i,
		}
		tp = append(tp, task)
	}
	return tp
}
