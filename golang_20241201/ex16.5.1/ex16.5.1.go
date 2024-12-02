package main

import "fmt"

/* 만약 슬라이스에 요소를 추가해야하는데, 남은 공간이 없다면? */

/* 방법 1: 함수 매개변수로 포인터를 받기(8byte) -> 단점 : 길이가 모자랄 경우 대비할 수 있음 */
// func addNum(slice *[]int) {
// 	*slice = append(*slice, 4)
// }

// func main() {
// 	slice := []int{1, 2, 3}
// 	addNum(&slice)

// 	fmt.Println(slice)
// }

/* 방법 2: 그냥 새로운 슬라이스를 반환하게 하기 (24byte)*/

func addNum(slice []int) []int {
	slice = append(slice, 4)
	// 함수 돌고 결과값으로 새로운 배열 반환하도록!
	return slice
}

func main() {
	slice := []int{1, 2, 3}
	slice = addNum((slice))

	fmt.Println(slice)
}
