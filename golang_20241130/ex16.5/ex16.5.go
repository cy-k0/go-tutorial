package main

import "fmt"

// tucker golang 개정판 챕터 반영

func main() {
	// append_test := []int{1, 2, 3, 4, 5}
	var append_test []int

	for i := 0; i <= 10; i++ {
		append_test = append(append_test, i)
	}

	fmt.Println(append_test)
	// append 함수는 기존 슬라이스 변경 X -> 새로운 슬라이스 반환
	// append_test2 := append(append_test, 6, 7, 8)

	// 기존 슬라이스의 값을 바꾸고싶다면?
	append_test = append(append_test, 6, 7, 8)
	fmt.Println(append_test)

	// fmt.Println(append_test2)
}
