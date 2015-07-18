package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
//	time.Sleep(5 * time.Second)
	say("hello")
	say("!")
}