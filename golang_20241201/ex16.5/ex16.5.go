package main

import (
	"fmt"
	"unsafe"
)

/*
string과 동일하게 슬라이스를 복사하면 구조체 필드 복사 (24byte 고정)
슬라이스의 길이가 길든 짧든, 결국 포인터 개념

type SliceHeader struct {
	Data uintptr -> 실제 배열을 가리키는 포인터
	Len int -> 요소 개수
	Cap int -> Capacity 최대 길이
	}
*/

func main() {
	slice1 := make([]int, 3, 5)

	slice2 := append(slice1, 4, 5, 6, 7, 8)

	slice1[0] = 100

	// 빈 공간이 남아있을땐, 기존 슬라이스에 append 하므로 같은 메모리 주소값을 가리킴
	// 빈 공간이 없을땐 새로운 슬라이스가 생성되므로 다른 메모리 주소를 기리킴
	// fmt.Printf("slice1: ", slice1, "slice1 pointer: %p\n", &slice1[0])
	// fmt.Printf("slice2: ", slice2, "slice2 pointer: %p\n", &slice2[0])
	fmt.Printf("slice1: ", slice1, "slice1 pointer: %v\n", unsafe.Pointer(&slice1[0]))
	fmt.Printf("slice2: ", slice2, "slice2 pointer: %v\n", unsafe.Pointer(&slice2[0]))
}
