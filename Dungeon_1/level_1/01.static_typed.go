package main

import "fmt"

// add function that takes two integers as arguments
func add(a int, b int) int {
    return a + b
}

func main() {
    // Correct usage
    // result := add(1, 2)
    // fmt.Println("Result:", result)

    //Incorrect usage - Uncommenting the following line will cause a compile-time error
    result = add(1, "two")
}