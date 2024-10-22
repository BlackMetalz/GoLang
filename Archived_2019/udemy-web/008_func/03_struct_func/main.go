package main

import "fmt"

type ahihi struct {
	a int
	b string
}

func (p ahihi) print_func() {
	fmt.Println("Hello world from print_func function")
	fmt.Println(p.b)
}

func main() {
	c := ahihi{
		a: 0,
		b: "Test",
	}

	c.print_func()
}
