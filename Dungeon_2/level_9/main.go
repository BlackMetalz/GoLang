// Looping with for loop
// print hello fucking world 3 times
package main

import "fmt"

func main(){
	for i := 1; i <= 5; i++ {
		// fmt.Printf("Lo Lo: %d time\n", i)
		if i == 2 {
			// fmt.Println("reached levle 2. Continue")
			continue
			// fmt.Println("reached levle 2. Break")
			//break 
		}
		fmt.Println(i)
	}

	k := 1
	for k <= 5 {
		fmt.Println(k*k)
		k++
	}
}