// Slice
// Continuous segment of an underlying array
// variable typed (elements can be added or removed!)
// More flexible
// Declare: slice := []int(1,2,3)
package main

import "fmt"

func main(){
	// Init slices
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	//           0 1 2 3 4 5 6 7 8 9
	fmt.Println(arr[1:8]) // [2 3 4 5 6 7 8]
	fmt.Println(arr[1:3]) // [2 3]
	fmt.Println(arr[:3]) // [1 2 3]
	fmt.Println(arr[5:]) // [6 7 8 9 10]

}