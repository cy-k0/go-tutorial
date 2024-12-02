package main

import "fmt"

func changeNum(array *[5]int) {
	array[4] = 100
}

func main() {
	array1 := [5]int{1, 2, 3, 4, 5}

	changeNum(&array1)

	slice1 := array1[1:2]
	fmt.Println("array1: ", array1)
	fmt.Println("slice: ", slice1)

}

/*
슬라이스 기본 : 시작 인덱스 포함, 끝 인덱스 미포함


처음부터 슬라이싱
slice[0:3] == slice[:3]

끝까지 슬라이싱
sliece[2:0] == slice[2:] == slice[2:len(slice)]

끝에서 하나 전까지 (python : slice[3:-1] --> golang 지원 안함)
slice[2:len(slice)-1]

**전체 슬라이싱
slice[:]

Q : 전체 슬라이싱 언제 쓸까?
배열을 슬라이스로 바꾸고 싶을때!
python은 배열을 복사하여 사용하나,
golang에서 슬라이스는 배열을 가리키고 있는 구조체 필드!
배열을 배열로 복사해오는 게 아니라, 배열 요소를 복사해 새로운 슬라이스를 생성해야함!


캡사이즈 조절 슬라이싱 (python은 이게 건너뛰는 인덱스인데..)
slice[시작인덱스 : 끝인덱스 : 최대인덱스]
--> 최대 인덱스 - 시작인덱스가 해당 슬라이스 구조체의 cap
*/
