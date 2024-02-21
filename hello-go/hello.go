package main

import (
	"fmt"
	"mat/rand"
	"time"
)

type T struct {
    name string // name of the object
    value int // its value
}

func main() {
	fmt.Println("Hello world!")

	fmt.Println("A hora agora é", time.Now())

	fmt.Println("My favorite number is", rand.Intn(10))
}