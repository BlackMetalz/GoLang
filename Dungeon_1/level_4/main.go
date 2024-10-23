package main

import "fmt"

func main() {
	var city string = "Hanoi"
	var country string = "Vietnam"
	var i int8 = 10
	fmt.Print("Hello from:", city,",",country, "\n")
	fmt.Print(city," ", country, "\n")
	fmt.Println("Hello from:", city,",",country)
	// printf
	fmt.Printf("Hello from: %s, %s\n", city, country)  // %s is a placeholder for strings
	fmt.Printf("Hello from: %v, %v\n", city, country) // %v is a placeholder for any value
	fmt.Printf("i: %d\n", i) // %d is a placeholder for integers
	fmt.Printf("i: %v\n", i)
}