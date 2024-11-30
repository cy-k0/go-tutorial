package main

import "fmt"

func main() {
	str1 := "죽는 날까지 하늘을 우러러 \n 한 점 부끄럼 없이"
	// 백쿼트로 묶으면 특수문자가 제대로 동작 x 여러 줄에 걸쳐 쓸 수 있음
	// '' 작은따움표는 문자 하나 나타낼때 씀
	str2 := `죽는 날까지 하늘을
	우러러 한 점 \n
	부끄럼 없이`

	fmt.Println(str1)
	fmt.Println(str2)
}
