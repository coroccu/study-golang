package main

import "fmt"

func add(x int, y int) {
	fmt.Println(x + y)
}

func addSub(x, y int) (int, int) {
	return x + y, x - y
}

func displayParams(params ...int) {
	fmt.Println(len(params), params)
}

func main() {
	add(1, 2)
	r1, r2 := addSub(10, 5)
	fmt.Println(r1, r2)

	f := func(x int) {
		fmt.Println("Internal Func1: ", x)
	}
	f(100)

	func(x int) {
		fmt.Println("Internal Func2: ", x)
	}(200)

	displayParams(10)
	displayParams(10, 20)
	displayParams(10, 20, 30)
	displayParams(10, 20, 30, 40)
	displayParams(10, 20, 30, 40, 50)

	ft := 1.11
	fmt.Println(int(ft))

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m, m)
}
