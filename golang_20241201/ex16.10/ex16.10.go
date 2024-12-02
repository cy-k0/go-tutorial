package main

import "fmt"

/* 중간 요소 제거 --> 인덱스 앞당겨주기 */

// for loop 사용
// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	idx := 2 // 제거할 인덱스

// 	for i := idx + 1; i < len(slice); i++ {
// 		slice[i-1] = slice[i]
// 	}

// 	slice = slice[:len(slice)-1] // 맨 끝 인덱스 제거

//		fmt.Println(slice)
//	}

// append 사용
// func main() {
// 	slice := []int{1, 2, 3, 4, 5}
// 	idx := 2

// 	slice = append(slice[:idx], slice[idx+1:]...)

// 	fmt.Println(slice)
// }

// copy 사용
func main() {
	slice := []int{1, 2, 3, 4, 5}
	idx := 2

	copy(slice[idx:], slice[idx+1:])
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
