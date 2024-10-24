package main

import "fmt"

func main() {
	var a,b string = "foo", "bar"
	fmt.Println(a + b) // foobar
	// increment
	var i int = 1
	i++
	fmt.Println(i) // 2
	// decrement
	i--
	fmt.Println(i) // 1
}