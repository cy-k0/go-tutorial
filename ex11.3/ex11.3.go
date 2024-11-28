package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("입력하세요")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("숫자를 입력하세요")
			// 개행문자 나올때까지 키보드 버퍼 읽어서 없애버리기
			stdin.ReadString('\n')
			fmt.Println()
			continue
		}
		fmt.Printf("입력하신 숫자는 %d 입니다", number)
		// 짝수 검사 하고 break
		if number%2 == 0 {
			break
		}
	}
}
