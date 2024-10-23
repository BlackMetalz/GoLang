// Using %T
package main

import "fmt"

func main() {
	var grades int = 42
	var msg string = "Hello, World!"

	fmt.Printf("grades is of type %T and msg is of type %T\n", grades, msg)
}