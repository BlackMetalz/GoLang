// if-else condition
package main

import "fmt"

func main() {

	var i int = 10
	if i <= 10 {
		fmt.Println("first")
	} else if i > 10 {
		fmt.Println("Second")
	} else {
		fmt.Println("WTF")
	}

	// Scan test
	var fruit string
	fmt.Println("Please enter your favourite fruit?")
	fmt.Scanf("%s", &fruit)
	if fruit == "apple" {
		fmt.Println("I love the fucking apple")
	} else if fruit == "banana" {
		fmt.Println("I love the fucking banana")
	} else {
		fmt.Println("I don't know what you want!")
	}
}