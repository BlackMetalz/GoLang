// Appending to a slice from slice xD
package main

import "fmt"

func main(){
	arr_1 := [5]int{10,20,30,40,50}
	slice_1 := arr_1[:2]

	arr_2 := [5]int{5,15,25,35,45}
	slice_2 := arr_2[:2]

	slice_3 := append(slice_1, slice_2...)
	fmt.Println(slice_3) // [10 20 5 15]


}