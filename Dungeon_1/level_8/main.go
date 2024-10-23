package main

import "fmt"

func main() {
	var i int // without give value to i, it will be 0. Called zero value
	var f float64
	fmt.Println("i:", i) 
	fmt.Printf("%d\n", i) 
	fmt.Printf("%.2f\n", f) // %.2f is a placeholder for float64 with 2 decimal points
	fmt.Printf("%f\n", f) // %f is a placeholder for float64 and print all decimal points
}