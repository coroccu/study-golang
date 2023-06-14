package main

import "fmt"

/* const */
const number = 1

const (
	name     = "Coroccu"
	password = "test"
)

func main() {
	/* var */

	/*
		var i int = 1
		var f float64 = 1.2
		var s string = "hoge"
		var t bool = true
	*/

	/*
		var (
			i int     = 1
			f float64 = 1.2
			s string  = "hoge"
			t bool    = true
		)
	*/

	i := 1
	f := 1.2
	s := "hoge"
	t := true

	fmt.Println(number, name, password)
	fmt.Printf("%T %T %T %T ", i, f, s, t)
	fmt.Println(i, f, s, t)
}
