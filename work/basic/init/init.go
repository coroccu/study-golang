package main

import "fmt"

func init() {
	fmt.Println("init")
}

func something() {
	fmt.Println("Do something!")
}

func main() {
	fmt.Println("init main")
	something()
}
