package main

import "fmt"


/*
const (
    Pi       = 3.14
    Language = "Go"
    Truth    = true
)
*/
func main() {
	// const Pi = 3.14
	const Greeting = "Hello, World!"
	Greeting = "Hello, Go!" // error
	// Pi = 3.141592222
    // fmt.Println("Pi:", Pi)
    fmt.Println(Greeting)
}

// ./main.go:16:2: cannot assign to Pi (neither addressable nor a map index expression)