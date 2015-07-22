package main
import (
	"time"
	"log"
	"os"
)

func panicingFunc() {
	panic("panicingFunc() panics")
}

func panicingFuncInnerCaller() {
	defer func() {
		log.Println("panicingFuncInnerCaller() ended")
	}()
	panicingFunc()
}

func panicGuard(action func()) {
	defer func() {
		if e := recover(); e != nil {
			log.Printf("PANIC: %v", e)
		}
	}()
	action()
}

func panicingFuncCaller() {
	defer func() {
		log.Println("panicingFuncCaller() ended")
	}()
	go panicGuard(panicingFuncInnerCaller)
	if _, err := os.Open("111.txt"); err != nil {
		log.Println(err)
	}
	time.Sleep(time.Second)
}

// created by missing full call stack. it should also contain main.main() G:/Dropbox/Go/src/github.com/nbgo/tour/panic-in-goroutines.go:39
//goroutine 5 [running]:
//main.panicingFunc()
//G:/Dropbox/Go/src/github.com/nbgo/tour/panic-in-goroutines.go:8 +0x6b
//main.panicingFuncInnerCaller()
//G:/Dropbox/Go/src/github.com/nbgo/tour/panic-in-goroutines.go:15 +0x3c
//created by main.panicingFuncCaller
//G:/Dropbox/Go/src/github.com/nbgo/tour/panic-in-goroutines.go:22 +0x46

func main() {
	defer func() {
		log.Println("Main() ended")
	}()
	panicingFuncCaller()
}