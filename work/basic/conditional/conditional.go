package main

import "fmt"

func main() {
	num := 4
	// num := 6
	// num := 10

	if num == 4 {
		fmt.Println("This is ", 4)
	} else if num == 6 {
		fmt.Println("This is ", 6)
	} else {
		fmt.Println("Others")
	}

	x, y := 10, 10
	// x, y := 10, 11

	if x == 10 && y == 10 {
		fmt.Println("And")
	}

	if x == 10 || y == 10 {
		fmt.Println("OR")
	}

	animal := "cat"
	switch animal {
	case "cat":
		fmt.Println("Cat!")
	case "dog":
		fmt.Println("Dog!")
	default:
		fmt.Println("Default!")
	}
}
