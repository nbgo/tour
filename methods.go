package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Float float64

func (f Float) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f *Float) AbsByRef() float64 {
	if *f < 0 {
		return float64(-(*f))
	}
	return float64(*f)
}


func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v.Abs(), v)

	f := Float(-math.Sqrt2)
	fmt.Println(f.Abs())
	fmt.Println((&f).AbsByRef())
}