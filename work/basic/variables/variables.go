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
	bt := true
	bf := false

	fmt.Println(number, name, password)
	fmt.Printf("%T %T %T %T %T ", i, f, s, bt, bf)
	fmt.Println(i, f, s, bt, bf)

	// cast
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	// array
	var a [2]int
	a[0] = 10
	a[1] = 20
	fmt.Println(a)

	//slice
	b := []int{1, 2, 3, 4}
	fmt.Println(b)
	fmt.Println(b[2])  //3
	fmt.Println(b[:2]) //1 2
	fmt.Println(b[:])  //all=1 2 3 4
	b[2] = 5
	fmt.Println(b) //1 2 5 4

	// var c = make([]int, 5)
	var c = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

	//map
	m := map[string]int{"Dogs": 100, "Cats": 200}
	fmt.Println(m)
	fmt.Println(m["Dogs"]) //100
	m["Pandas"] = 300
	fmt.Println(m)
	fmt.Println(m["none"]) //0

	//byte
	byt := []byte{65, 66}
	fmt.Println(byt)
	fmt.Println(string(byt)) //AB ASCII Codes
}
