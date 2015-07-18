package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v)
	v.X = 4
	fmt.Println(v)

	p := &v
	p1 := &v.Y
	p.X = 1e9
	*p1 = 5
	fmt.Println(v)

	v2 := Vertex{X:1}
	p2 := &Vertex{1, 2}
	fmt.Println(v2, p2)
}