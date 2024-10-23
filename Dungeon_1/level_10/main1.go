// using reflect.TypeOf() to get the type of a variable
package main

import (
	"fmt"
	"reflect"
)

func main() {

	var grades int = 42
	var msg string = "Hello, World!"

	fmt.Println("grades is of type", reflect.TypeOf(grades), "and msg is of type", reflect.TypeOf(msg))
}