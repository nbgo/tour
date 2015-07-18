package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (this Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, this)
}

func (this String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, this)
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}