package main

import "fmt"

var city string = "HaLong" // This is also called global variable

func main() {
	city := "Hanoi" // This is also called local variable
	{
		country := "Vietnam"
		fmt.Println("Hello from:", city, ",", country)
	}

	// Outer block can not access the variable declared in the inner block. country is in the inner block
	// fmt.Println("Hello from:", city, ",", country) // This line will cause a compile-time error


}