package main

import (
	"fmt"
	"time"
)

func fibonacci(c chan int) {
	x, y, n := 0, 1, cap(c)
	for i := 0; i < n; i++ {
		fmt.Println("step", i)
		c <- x
		x, y = y, x+y
		time.Sleep(100 * time.Millisecond)
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(c)
	for i := range c {
		fmt.Println(i)
	}
}