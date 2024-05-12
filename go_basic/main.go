package main

import "fmt"

func double(n *int) {
	*n *= 2
}

func main() {
	// b := new(int)
	// c := new(*int)

	// a := 1
	// b = &a
	// c = &b

	// fmt.Printf("address a %v\n", &a)
	// fmt.Printf("address b %v\n", &b)
	// fmt.Printf("address c %v\n", &c)

	// fmt.Println()

	// fmt.Printf("value a %v\n", a)
	// fmt.Printf("value b %v\n", b)
	// fmt.Printf("value c %v\n", c)

	// fmt.Println()

	// fmt.Printf("*b %v\n", *b)
	// fmt.Printf("*c %v\n", *c)
	// fmt.Printf("**c %v\n", **c)

	// fmt.Println()

	// x := 2
	// var n *int = &x

	// println(x)
	// println(n)

	// double(n)

	// fmt.Println(*n)
	// fmt.Println(x)

	// fmt.Println()

	y := 2

	println(y)

	double(&y)

	fmt.Println(y)
}
