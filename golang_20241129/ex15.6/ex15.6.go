package main

import "fmt"

func main() {
	var char rune = '한'
	fmt.Printf("%T\n", char)
	// rune 타입은 int32 -> 숫자값이 반환 됨
	fmt.Println(char)
	// c 포멧팅 사용하여 출력 가능
	fmt.Printf("%c", char)
}
