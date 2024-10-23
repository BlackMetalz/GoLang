package main

import "fmt"

func main() {
	var door int
	var is_open bool
	fmt.Print("You are in the dungeon. There are 3 doors in front of you. Which one do you choose? (1, 2, 3): ")
	count, err := fmt.Scanf("%d %t", &door, &is_open)
	fmt.Println("count:", count, "error:", err)
	fmt.Println("You chose door", door, "and it is open:", is_open)
}