package main

import "fmt"

func main() {
	// var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	// arr1 := [...]int{1, 2, 3, 4, 5} // 뒤에 오는 값 개수에 따라 배열 길이가 정해짐 (고정)
	// var arr1 = [5]int{1: 10, 3: 30} 1번 index에 10, 3번 index에 30, 나머지는 기본값 0
	arr1 := []int{1, 2, 3, 4, 5} // 길이가 늘어나거나 줄어들 수 있는 동적 배열
	for i := 0; i < 5; i++ {
		fmt.Println(arr1[i])
	}
}
