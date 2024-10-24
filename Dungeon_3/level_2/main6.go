// Deleting from a slice
package main

import "fmt"

func main(){
	arr := [5]int{10,20,30,40,50}
	// task: delete 30 ( index 2 ) from arr 
	i := 2
	s1 := arr[:i]
	s2 := arr[i+1:]
	s3 := append(s1,s2...)
	// With ...
	// Efficiency: More efficient as it appends all elements in one call.
	// Readability: More concise and easier to read
	// Without ...
	/*
    // Append each element of s2 to s1 individually
    s3 := append(s1, s2[0])
    s3 = append(s3, s2[1])
    s3 = append(s3, s2[2])
	*/
	fmt.Println(s3) // [10 20 40 50]
}