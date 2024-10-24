// init slice #2
package main

import "fmt"

func main(){
	arr := [10]int{10,20,30,40,50,60,70,80,90,100}
	slice := arr[1:8]
    fmt.Println(cap(arr))   // Output: 10
    fmt.Println(slice)      // Output: [20 30 40 50 60 70 80]
    fmt.Println(cap(slice)) // Output: 9
	// Why capacity of slice is 9? cap of arr - start index of slice
	// Other
	slice1 := arr[5:6]
	fmt.Println(slice1) // [60]
	fmt.Println(cap(slice1)) // 10-5 = 5?
}
