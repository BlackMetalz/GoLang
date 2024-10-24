// Arrays
// Declare
// var <array name> <size of the array> <data type>
// var grades [5] int
// var fruit [3] string
package main

import "fmt"

func main(){
	var grades [5]int = [5]int{1,2,3,4,5}
	fmt.Println(grades)
	fmt.Println(grades[0]) //First index
    grades[0] = 99 // change the value of index
	fmt.Println(grades) 
	fmt.Println(grades[0]) //First index
	var fruit [3]string = [...]string{"banana", "apple", "orange"}
	fmt.Println(len(fruit))
	marks := [3]int{10,20,30}
	fmt.Println(marks)
	// loop through an array with length
	fmt.Println("=============")
	for i := 0; i < len(grades); i++ {
		fmt.Println(grades[i])
	}
	fmt.Println("=============")
	// Looping through an array
	for index, element := range grades {
		fmt.Println(index, "=>", element)
	}
	fmt.Println("=============")
	// Multi dimensional arrays
	arr := [3][2]int{ {1,2}, {3,4}, {5,6}}
	fmt.Println(arr[0][1]) // 2
	fmt.Println(arr[1][0]) // 3
	fmt.Println(arr[2][1]) // 6

}