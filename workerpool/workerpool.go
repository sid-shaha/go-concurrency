package main

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	NumberOfWorkers int
	Tasks           []TaskProcessor
	wg              sync.WaitGroup
	inputChannel    chan TaskProcessor
}

func (wp *WorkerPool) Worker(id int) {
	defer wp.wg.Done()

	for task := range wp.inputChannel {
		fmt.Println("Task taken by the worker with id", id)
		task.DoTask()
	}

	fmt.Println("Worker killed as tasks are done", id)
}
func (wp *WorkerPool) Run() {
	wp.inputChannel = make(chan TaskProcessor, len(wp.Tasks))

	wp.wg.Add(wp.NumberOfWorkers)

	for i := 0; i < wp.NumberOfWorkers; i++ {
		go wp.Worker(i)
	}

	for i := 0; i < len(wp.Tasks); i++ {
		wp.inputChannel <- wp.Tasks[i]
	}
	close(wp.inputChannel)

	wp.wg.Wait()

	fmt.Println("Finished with all the processing")

}
