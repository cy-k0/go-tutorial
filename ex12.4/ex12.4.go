package main

import "fmt"

func main() {
	x := [...]float64{10.1, 20.2, 30.3, 40.4, 50.5}
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	for _, v := range x {
		fmt.Println(v)
	}
}
