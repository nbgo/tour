package main
import (
	"fmt"
	"time"
)

func main() {
	case2()
	fmt.Println("end")
}

func case1() {
	ch := make(chan int, 1) // without buffered channel we'll get deadlock after goroutine ends
	go func() {
		time.Sleep(2*time.Second)
	}()
	ch <- 1
}

func case2() {
	ch := make(chan int, 1)
	for i := 0; i<5; i++ {
		go func(i int) {
			// ch <- doQuery(i+1) // it will block until full channel is read

			select {
			case ch <- doQuery(i+1):
				fmt.Println("result sent to reader or buffer") // it will write if there is a reader or can write if there is a buffer
			default:
				fmt.Println("no reader and buffer")
			}
			fmt.Println("goroute exit")
		}(i)
	}
	time.Sleep(2*time.Second)
	fmt.Println(<-ch) // deadlock will be here without buffered channel
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
	time.Sleep(time.Second)
}

func doQuery(t int) int {
	fmt.Println("query started")
	time.Sleep(time.Duration(t) * time.Millisecond)
	fmt.Println("query ended")
	return t
}