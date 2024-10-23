package main
import "fmt"

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello, %s!\n", name)
	fmt.Print("You are in the dungeon. There are 3 doors in front of you. Which one do you choose? (1, 2, 3): ")
	var door int
	var is_open bool
	fmt.Scanf("%d %t", &door, &is_open)
	fmt.Println("You chose door", door, "and it is open:", is_open)

}