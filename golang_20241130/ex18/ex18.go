package main

import "fmt"

func main() {
	int_arr := []int{11, 12, 13}

	// for i := 0; i < len(int_arr); i++ {
	// 	int_arr[i] += 10
	// }

	// range 를 사용하면 index가 함께 반환됨을 기억
	for i, v := range int_arr {
		int_arr[i] = v + 10
	}
	fmt.Println(int_arr)
}
