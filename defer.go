package main

import "fmt"

func main() {
	fmt.Println("counting")

	defer fmt.Println("first")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("second")

	fmt.Println("done")
}