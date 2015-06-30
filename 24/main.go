package main

import (
	"fmt"
	"time"
)

var counter = 0

func main() {
	for i := 0; i < 2; i++ {
		go incr()
	}
	time.Sleep(time.Millisecond * 10)
}

func incr() {
	counter++
	fmt.Println(counter)
}
