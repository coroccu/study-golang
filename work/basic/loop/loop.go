package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("============")
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("Continued!")
			continue
		}

		if i > 7 {
			fmt.Println("Broke!")
			break
		}
		fmt.Println(i)
	}
	fmt.Println("============")
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

	// infinite loop
	// for {
	// 	fmt.Println("infinite loop")
	// }

	// range
	arrayAnimal := []string{"cat", "dog", "pig"}

	for i := 0; i < len(arrayAnimal); i++ {
		fmt.Println(i, arrayAnimal[i])
	}

	for i, v := range arrayAnimal {
		fmt.Println(i, v)
	}

	for _, v := range arrayAnimal {
		fmt.Println(v)
	}
}
