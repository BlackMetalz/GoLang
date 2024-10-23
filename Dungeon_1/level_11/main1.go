// strconv package
// The strconv package provides functions for converting strings to other types and vice versa.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "42"
	var i int = 33
	var num int
	var s string = strconv.Itoa(i)
	num, _ = strconv.Atoi(str)
	
	fmt.Println(num)
	fmt.Println(s)
	fmt.Printf("%q\n", s) // return quoted string "33"
	fmt.Printf("%q\n", strconv.Itoa(i)) // return quoted string "33"
	// String to int
	// The Atoi function converts a string to an integer.
	// If the string is not a valid integer, it returns an error.
	// The second return value is an error value.
	// If the conversion fails, the error value is not nil.
	// If the conversion succeeds, the error value is nil.
	// The underscore (_) is used to ignore the error value.
}