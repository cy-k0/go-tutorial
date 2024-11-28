package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	b := [...]int{100, 200, 300, 400, 500}

	fmt.Println()
	for i, v := range a {
		fmt.Println("a[%d] = %d\n", i, v)
	}
	fmt.Println()

	for i, v := range b {
		fmt.Println("b[%d] = %d\n", i, v)
	}
	b = a
	fmt.Println()
	for i, v := range b {
		fmt.Println("b[%d] = %d\n", i, v)
	}

}
