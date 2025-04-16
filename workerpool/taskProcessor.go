package main

import (
	"fmt"
	"math/rand"
	"time"
)

type TaskProcessor interface {
	DoTask()
}

type Image struct {
	Id int
}

func (ig *Image) DoTask() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering from the panic, task killed with id", ig.Id)
		}
	}()
	fmt.Println("Started processing image with id ", ig.Id)
	if rand.Intn(5) == 0 {
		panic("Random panic triggered!")
	}
	time.Sleep(time.Second * 5)
	fmt.Println("Done processing image with id ", ig.Id)
}
