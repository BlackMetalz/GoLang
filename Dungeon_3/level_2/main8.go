// Looping through a slice is same as looping through array

package main

import "fmt"

func main(){
	arr := []int{10,20,30,40,50}

	for key,value := range arr {
		fmt.Println("Key:", key, "And Value: ", value)
	}

	// In case we don't need the key/index
	for _, value := range arr {
		fmt.Println("value only: ", value)
	}
}
