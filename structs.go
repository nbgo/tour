package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var globalVertex *Vertex
func Test1() Vertex {
	globalVertex = &Vertex{1,1}
	return *globalVertex
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

	v3 := Test1()
	fmt.Println(globalVertex)
	v3.X = 5
	fmt.Println(globalVertex)

	sl := make([]Vertex, 1)
	v4 := Vertex{1,1}
	sl[0] = v4
	v4.X = 2
	fmt.Println(sl[0]) // output is 1,1

	sl2 := make([]*Vertex, 1)
	v5 := &Vertex{1,1}
	sl2[0] = v5
	v5.X = 2
	fmt.Println(sl2[0]) // output is 2,1
}