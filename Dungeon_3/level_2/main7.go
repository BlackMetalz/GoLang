// Copying from a slice
package main

import "fmt"

func main(){
	arr := []int{10,20,30,40,50}
	dest := make([]int, 3)
	num := copy(dest, arr)
	fmt.Println(dest)
	fmt.Println("Number of slice copied: ", num, "and the first element is: ", dest[0])
}