package main

import (
	"fmt"
	"sort"
)

/* 구조체 슬라이스 정렬 */
type Student struct {
	Name string
	Age  int
}

// 별칭타입
type Students []Student

// Sort 를 사용하기 위해 메서드 생성
func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 18},
		{"백두산", 12},
		{"하나", 33},
	}
	sort.Sort(Students(s))
	fmt.Println(s)

}
