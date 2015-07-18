package main

import (
	"fmt"
	"math/cmplx"
)

var c, python, java bool
var autoBool, autoInt, autoString, autoFloat = true, 1, "no!", 3.1415962

var (
	toBe = false
	maxInt uint64 = 1<<64-1
	z = cmplx.Sqrt(-5 + 12i)
)

func main() {
	j := 3
	var i = 2
	const f = "%T(%v)\n"
	fmt.Println(i, c, python, java)
	fmt.Println(autoBool, autoInt, autoString, autoFloat)
	fmt.Printf("%q\n", autoString)
	fmt.Println(j)
	fmt.Printf(f, toBe, toBe)
	fmt.Printf(f, maxInt, maxInt)
	fmt.Printf(f, z, z)
}