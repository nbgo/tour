package main

import (
	"fmt"
	"strings"
	"reflect"
)


type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func WordCount(s string) (result map[string]int) {
	result = make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		result[word]++
	}
	return result
}

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)

	m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(m)

	m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google": {37.42202, -122.08408},
	}
	fmt.Println(m)

	fmt.Println(WordCount("I am am learning Go!"))
	fmt.Printf("Is result correct? %v", reflect.DeepEqual(map[string]int {"am": 2, "learning": 1, "Go!": 1, "I": 1}, WordCount("I am am learning Go!")))
}