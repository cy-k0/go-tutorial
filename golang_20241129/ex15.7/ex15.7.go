package main

import "fmt"

// func main() {
// 	var str string = "Hello 월드"

// 	// len 은 문자 개수가 아니라 byte 길이를 반환
// 	for i := 0; i < len(str); i++ {
// 		fmt.Printf("타입: %T, 값: %d, 문자: %c\n", str[i], str[i], str[i])
// 	}
// }

/*
12 번 순회, str[i] 는 인덱스 값이 아니라, byte 자리
타입: uint8, 값: 72, 문자: H
타입: uint8, 값: 101, 문자: e
타입: uint8, 값: 108, 문자: l
타입: uint8, 값: 108, 문자: l
타입: uint8, 값: 111, 문자: o
타입: uint8, 값: 32, 문자:
타입: uint8, 값: 236, 문자: ì
타입: uint8, 값: 155, 문자:
입: uint8, 값: 148, 문자:
타입: uint8, 값: 235, 문자: ë
타입: uint8, 값: 147, 문자:
타입: uint8, 값: 156, 문자:
*/

// 배열 순회를 사용해 인덱스를 사용할 수 있다
// 타입변환+len, range 두가지 사용
func main() {
	var str string = "Hello 월드"
	// var arr []rune = []rune(str)
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Printf("타입: %T, 값: %d, 문자: %c\n", arr[i], arr[i], arr[i])
	for _, v := range str {
		fmt.Printf("타입: %T, 값: %d, 문자: %c\n", v, v, v)
	}
}

/*
타입: int32, 값: 72, 문자: H
타입: int32, 값: 101, 문자: e
타입: int32, 값: 108, 문자: l
타입: int32, 값: 108, 문자: l
타입: int32, 값: 111, 문자: o
타입: int32, 값: 32, 문자:
타입: int32, 값: 50900, 문자: 월
타입: int32, 값: 46300, 문자: 드
*/
