package main

import "fmt"

/*
슬라이스 복제
언제 슬라이스 복제가 필요할까

슬라이스의 값을 복제해 오되, 일부 요소 값을 바꾸고싶을때!

왜냐하면 슬라이스는 구조체! 포인터, 길이, 최대길이를 가진 구조체!
그래서 복제해왔을때 같은 메모리 공간을 바라보고, 해당 데이터를 직접 바꿔버림!
이게 단점이 되기도 한다는 것!
slice1의 값은 그대로 두고, slice2에서 slice1을 복제해 온 뒤 일부 요소만 바꾼 뒤
slice1을 출력해보면 함께 바뀌거등요
*/

// for loop 이용한 방법
// func main() {
// 	slice1 := []int{1, 2, 3, 4, 5, 6}

// 	slice2 := make([]int, len(slice1))

// 	for i, v := range slice1 {
// 		slice2[i] = v
// 	}
// 	slice2[1] = 100

// 	fmt.Println(slice1)
// 	fmt.Println(slice2)
// }

// append를 이용한 방법 (for loop 동일)
// func main() {
// 	slice1 := []int{1, 2, 3, 4, 5}
// 	slice2 := append([]int{}, slice1...)

// 	slice2[1] = 100
// 	fmt.Println(slice1)
// 	fmt.Println(slice2)
// }

// copy 를 이용한 방법
func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	// 소스에서 목적지로! copy(dst, src) --> 가독성 좋음
	copy(slice2, slice1)
	slice2[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)
}
