package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s =", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %d\n", i, s[i])
	}

	fmt.Println("s[1:4] ==", s[1:4])
	fmt.Println("s[:3] ==", s[:3])
	fmt.Println("s[4:] ==", s[4:])
	fmt.Println("s[4:len(s)] ==", s[4:len(s)])
}