// Append to a slice
package main

import "fmt"

func main(){
	// slice = append(slice, ele1,ele2)
	arr := [4]int{10,20,30,40}

	slice := arr[1:3] // [20,30]
	fmt.Println(slice)
	fmt.Println(len(slice)) // 2
	fmt.Println(cap(slice)) // 3

	slice = append(slice, 9000,-90,99)
	fmt.Println(slice)
	fmt.Println(len(slice)) // 5
	fmt.Println(cap(slice)) // 6 ( this shit was double )

	slice = append(slice, 111,-22)
	fmt.Println(slice)
	fmt.Println(len(slice)) // 7
	fmt.Println(cap(slice)) // 12 ( this shit was double from 6 )
}