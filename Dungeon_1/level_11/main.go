// Convert int to float64

package main

import "fmt"
import "reflect"

func main() {
	var x int = 42
	var y float64 = float64(x)
	fmt.Println(y)
	fmt.Printf("%.2f\n", y)
	fmt.Println("type of y:", reflect.TypeOf(y))
}