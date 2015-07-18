package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	// can't use v here, though
	return lim
}

func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func Sqrt(x float64) (float64, int) {
	var (
		z = 1.0
		i int
	)
	for i = 0;; i++ {
		prevZ := z
		z -= (z * z - x) / (2 * z)
		if math.Abs(prevZ - z) < 0.00000000001 {
			break
		}
	}
	return z, i
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
	fmt.Println(Sqrt(4))
	fmt.Println(math.Sqrt(4))
	fmt.Println(Sqrt(10))
	fmt.Println(math.Sqrt(10))
}