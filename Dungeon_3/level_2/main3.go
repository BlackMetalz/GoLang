// Slice and index numbers
package main

import "fmt"

func main(){
	arr := [5]int{10,20,30,40,50}
	slice := arr[:3]
	fmt.Println(slice) // [10 20 30]
	// fmt.Println(slice[1])
	slice[1] = 9000
	fmt.Println("After modification")
	fmt.Println(arr) // [10 9000 30 40 50]
	fmt.Println(slice) // [10 9000 30]
}