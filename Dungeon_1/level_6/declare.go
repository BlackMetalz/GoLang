package main

import "fmt"

func main() {
  // Declare and initialize variables
  var t,s string = "Hello", "World"
  fmt.Println(t,s)

  // Declare and initialize variables with different data types
  var (
	s1 string = "Hong"
	i int = 10
  )
  fmt.Println(s1,i)

  // Short variable declaration
  s2 := "Kong"
  i2 := 20
  fmt.Println(s2,i2)

  s3 := "peter"
  s3 = "parker"
  fmt.Println(s3)
}