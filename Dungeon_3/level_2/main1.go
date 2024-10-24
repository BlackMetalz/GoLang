// Init slice way
package main

import "fmt"

func main(){
	slice := make([]int, 5,8)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}