package main

import (
	"fmt"
	"sort"
)

/*
sort 패키지를 활용한 정렬
***주의!***
sort 패키지는 Interface 타입을 인자로 받음
(len, less, swap 을 가진)
*/
func main() {
	slice := []int{100, 8, 44, 77, 20}
	sort.Ints(slice)

	fmt.Println(slice)

}
