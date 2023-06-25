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

	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	min := l[0]
	for i := 1; i < len(l); i++ {
		if l[i] < min {
			min = l[i]
		}
	}
	fmt.Println(min)

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	sum := 0
	for _, v := range m {
		sum += v
	}
	fmt.Println(sum)
}
