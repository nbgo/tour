package main

//import "golang.org/x/tour/pic"
import "fmt"

func Pic(dx, dy int) (result [][]uint8) {
	result = make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		result[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			result[y][x] = uint8(x^y)// uint8((x + y) / 2)
		}
	}
	return
}

func main() {
	//pic.Show(Pic)
	fmt.Println(Pic(10, 10))
}