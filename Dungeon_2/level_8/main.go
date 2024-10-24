// Switch case
package main

import "fmt"

func main(){
	var i int = 10
	switch i {
	case 1:
		fmt.Println("case 1")
		fallthrough
	case 10:
		fmt.Println("case 10")
		// fallthrough
	case 20:
		fmt.Println("case 20")
		fallthrough
	default:
		fmt.Println("case default")
	}

	// switch with conditions
	var a,b int = 10,20
	switch {
	case a+b == 30:
		fmt.Println("a+b=30")
	case a+b >30:
		fmt.Println("a+b>30")
	default:
		fmt.Println("default case")

	}


}