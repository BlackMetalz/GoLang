package main

import "fmt"

func main() {
    a := 42
    b := 3.14159
    c := "Hello, World!"
    d := true
	zz := "zz"
    fmt.Printf("Default format: %v\n", a)
    fmt.Printf("Type of value: %T\n", a)
    fmt.Printf("Decimal integer: %d\n", a)
    fmt.Printf("Binary representation: %b\n", a)
    fmt.Printf("Octal representation: %o\n", a)
    fmt.Printf("Hexadecimal representation (lowercase): %x\n", a)
    fmt.Printf("Hexadecimal representation (uppercase): %X\n", a)
    fmt.Printf("Floating-point number: %f\n", b)
    fmt.Printf("Scientific notation (lowercase): %e\n", b)
    fmt.Printf("Scientific notation (uppercase): %E\n", b)
    fmt.Printf("String: %s\n", c)
    fmt.Printf("Quoted string: %q\n", c)
    fmt.Printf("Pointer address: %p\n", &a)
    fmt.Printf("Boolean: %t\n", d)
	fmt.Printf("String: %c\n", zz) // wrong type of format
}