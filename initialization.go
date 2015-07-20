package main
import "fmt"

func main() {
	v1 := [2]int{2, 3} // array
	v2 := [...]int{2, 3, 5, 7, 11, 13} // array
	var v3 [3]int // array
	v4 := []int{2, 3, 5, 7, 11, 13} // slice
	v5 := v2[:] // slice
	v6 := v2[:] // slice referencing the same array as v2
	v7 := v1 // new array copied from v1
	v8 := new([5]int)
	var v9 []int // slice equals nil
	v10 := [0]int{} // array of zero size
	v11 := make([]int, 10) // slice


	fmt.Printf("Type of v1 = %T\n", v1)
	fmt.Printf("Type of v2 = %T\n", v2)
	fmt.Printf("Type of v3 = %T\n", v3)
	fmt.Printf("Type of v4 = %T\n", v4)
	fmt.Printf("Type of v5 = %T\n", v5)
	v5[0] = 1
	fmt.Printf("v1[0]=%v, v7[0]=%v\n", v5[0], v6[0])
	v1[0] = 0
	fmt.Printf("v1[0]=%v, v7[0]=%v\n", v1[0], v7[0])
	fmt.Printf("v1 == v1: %v\n", v1 == v1)
	fmt.Printf("v1 == v7: %v\n", v1 == v7)
	fmt.Printf("Type of v8 = %T\n", v8)
	fmt.Printf("v9 == nil: %v\n", v9 == nil)
	fmt.Printf("Type of v10 = %T\n", v10)
	fmt.Printf("Type of v11 = %T\n", v11)
}
