package main

import "fmt"

func hello(name string) {
	fmt.Println("Hello", name)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	if b == 0 {
		fmt.Println("Error: Division by zero.")
		return 0
	}
	return a / b
}

func main() {
	name := "Dot8Pixels"
	hello(name)

	// lambda function
	func(name string) {
		fmt.Println("Hello", name)
	}(name)

	result := func(a int, b int) int {
		return a + b
	}(1, 2)

	fmt.Println(result)

	fmt.Println(add(4, 2))
	fmt.Println(sub(4, 2))
	fmt.Println(mul(4, 2))
	fmt.Println(div(4, 2))
}
