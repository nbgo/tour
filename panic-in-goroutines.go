package main
import (
	"fmt"
	"time"
)

func panicingFunc() {
	panic("panicingFunc() panics")
}

func panicingFuncInnerCaller() {
	defer func() {
		fmt.Println("panicingFuncInnerCaller() ended")
	}()
	panicingFunc()
}

func panicingFuncCaller() {
	defer func() {
		fmt.Println("panicingFuncCaller() ended")
	}()
	go panicingFuncInnerCaller()
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
		fmt.Println("Main() ended")
	}()
	panicingFuncCaller()
}