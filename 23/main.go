package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go process()
	time.Sleep(time.Millisecond * 10) // Bad!
	fmt.Println("done")

	go func() {
		fmt.Println("procesing2")
	}()
	// Won't work if we don't sleep
	time.Sleep(time.Millisecond * 10) // Bad!
}

func process() {
	fmt.Println("processing")
}
