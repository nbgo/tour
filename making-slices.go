package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:3]
	printSlice("d", d)

	var z []int
	printSlice("z", z)
	if z == nil {
		fmt.Println("nil!")
	}

	z = append(z, 0)
	printSlice("z after 1st append", z)

	z = append(z, 1)
	printSlice("z after 2nd append", z)

	z = append(z, 2)
	printSlice("z after 3rd append", z)

	// Slicing does not copy the slice's data. It creates a new slice value that points to the original array
	z1 := z[0:1];
	z1[0] = 3;
	printSlice("z", z)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}