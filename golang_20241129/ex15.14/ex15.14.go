package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	Data uintptr
	Len  int
}

func main() {
	str1 := "Hello 월드"
	str2 := str1

	stringHeader1 := (*StringHeader)(unsafe.Pointer(&str1))
	stringHeader2 := (*StringHeader)(unsafe.Pointer(&str2))

	fmt.Printf("str1 stringHeader: %d\n", stringHeader1)
	fmt.Printf("str2 stringHeader: %d", stringHeader2)
}
