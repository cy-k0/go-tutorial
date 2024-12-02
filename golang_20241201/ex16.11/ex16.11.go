package main

import "fmt"

/* 요소 삽입 */
//for loop 사용
// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

// 	slice = append(slice, 0) // 삽입하고싶은 값
// 	idx := 2                 // 삽입하고싶은 위치

// 	// 요소값이 하나 추가됬으므로 len(slice) = 10
// 	for i := len(slice) - 2; i >= idx; i-- {
// 		slice[i+1] = slice[i]
// 	}

// 	slice[3] = 100
// 	fmt.Println(slice)
// }

// append 사용 --> 메모리 효율 떨어짐
// func main() {
// 	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	idx := 2

// 	slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)

// 	fmt.Println(slice)
// }

// copy 사용 --> 불필요한 메모리 할당 없음
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	idx := 2 // 삽입하고싶은 위치

	slice = append(slice, 0)
	copy(slice[idx+1:], slice[idx:])
	slice[3] = 100
	fmt.Println(slice)
}
