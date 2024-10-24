// operator, arithmethic

package main

import "fmt"

func main() {
    a := 10
    b := 3

    // Addition
    sum := a + b
    fmt.Printf("%d + %d = %d\n", a, b, sum)

    // Subtraction
    diff := a - b
    fmt.Printf("%d - %d = %d\n", a, b, diff)

    // Multiplication
    prod := a * b
    fmt.Printf("%d * %d = %d\n", a, b, prod)

    // Division
    quot := a / b
    fmt.Printf("%d / %d = %d\n", a, b, quot)

    // Modulus
    mod := a % b
    fmt.Printf("%d %% %d = %d\n", a, b, mod)
}