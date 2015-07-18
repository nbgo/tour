package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	var step, val1, val2 = 0, 0, 1
	return func() int {
		step++
		switch step {
		case 1: return val1
		case 2: return val2
		default:
			curr := val1 + val2
			val1, val2 = val2, curr
			return curr
		}
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}