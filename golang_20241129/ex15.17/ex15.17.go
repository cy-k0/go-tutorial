package main

import (
	"fmt"
	"strings"
)

// 합 연산 사용 : go 언어 내부에서는 합 연산을 사용할때마다 새로운 메모리 공간을 할당
// 메모리 공간이 버려짐
func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		// c가 소문자면, 소문자 a랑 얼마나 떨어져있는가
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a'))
		} else {
			rst += string(c)
		}
	}
	return rst
}

// strings.Builder 는 내부에 슬라이스를 가지고 있으므로 WriteRune() 메서드를 통해 문자를 더할때 메모리를 새로 생성하지 않음.
// 기존 메모리 공간에 빈자리가 있으면 그냥 더함 -> 메모리 공간 낭비 없앨 수 있음
func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a'))
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func main() {
	var str string = "hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}
