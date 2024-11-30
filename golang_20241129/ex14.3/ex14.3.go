package main

import "fmt"

type Data struct {
	value int
	data  [1]int
}

// 함수의 인자는 무조건 r value (값으로 쓰인다)
func ChangeData(arg *Data) {
	arg.value = 999
	arg.data[0] = 999
}

func main() {
	var data Data
	// newData := Data{123, [...]int{100}}

	// 값 복사 방식이므로, 함수 내부의 arg는 호출 시 전달된 data 변수의 복사본
	ChangeData(&data)

	// data = newData

	fmt.Println(data.value)
}
