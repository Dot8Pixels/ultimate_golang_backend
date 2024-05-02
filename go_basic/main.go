package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	var sum int
	for i := 1; i <= 10; i++ {
		sum += 1
		fmt.Println(sum)
	}

	for i := range 3 {
		fmt.Println(i + 1)
	}
}
