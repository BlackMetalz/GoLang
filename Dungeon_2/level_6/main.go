package main

import "fmt"

func main() {
	// & operator
	var x,y int = 12,25
	z := x & y
	fmt.Println(z) // 8
	// | operator
	z = x | y
	fmt.Println(z) // 29
	// ^ operator
	z = x ^ y
	fmt.Println(z) // 21
	// << operator
	z = x << 1
	fmt.Println(z) // 24
	// >> operator
	z = x >> 1
	fmt.Println(z) // 6
}