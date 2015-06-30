package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func main() {
	// Channels have native buffers
	//c := make(chan int, 100)
	c := make(chan int)
	for i := 0; i < 4; i++ {
		worker := Worker{id: i}
		// Make some workers which will idle
		// until they have work
		go worker.process(c)
	}

	//for {
	//// Dropping
	//select {
	//case c <- rand.Int():
	//// Optional
	//default:
	//fmt.Println("dropped")
	//}
	//time.Sleep(time.Millisecond * 50)
	//}

	//for {
	//select {
	//case c <- rand.Int():
	//case <-time.After(time.Millisecond * 100):
	//fmt.Println("timed out")
	//}

	//time.Sleep(time.Millisecond * 50)
	//}

	for {
		select {
		case c <- rand.Int():
		case t := <-time.After(time.Millisecond * 100):
			fmt.Println("timed out", t)
		default:
			fmt.Println("no channel")
		}

		time.Sleep(time.Millisecond * 50)
	}
}

func after(d time.Duration) chan bool {
	c := make(chan bool)
	go func() {
		time.Sleep(d)
		c <- true
	}()
	return c
}

func (w Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}
