package main

import (
	"fmt"
	"math"
	"reflect"
)

type Abser interface {
	Abs() float64
}

type NagativeAbser interface {
	NegativeAbs() float64
}

type FullAbser interface {
	Abser
	NagativeAbser
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
//	a = v

	fmt.Println(a.Abs())

	fmt.Println(v.Abs())
	fmt.Println(v.NegativeAbs())
	fmt.Println(v)
	fmt.Println(&v)

	var a1 FullAbser
	a1 = &v
	fmt.Println(a1.Abs())
	fmt.Println(a1.NegativeAbs())
	fmt.Println(a1)

	var c *Vertex = nil
	var b Abser = c
	fmt.Println(reflect.TypeOf(nil))
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(b == nil)
	fmt.Println(reflect.ValueOf(b).IsNil())
	fmt.Println(c == nil)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) NegativeAbs() float64 {
	return -v.Abs()
}

func (v Vertex) String() string {
	return fmt.Sprintf("X=%v Y=%v", v.X, v.Y)
}

//func (v Vertex) String() string {
//	return fmt.Sprintf("X=%v Y=%v", v.X, v.Y)
//}
