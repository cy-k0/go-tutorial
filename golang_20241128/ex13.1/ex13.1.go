package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	// var myHouse House = House{
	// 	"서울시 강동구",
	// 	24,
	// 	9.80,
	// 	"아파트",
	// }
	// myHouse.Address = "서울시 강남구"
	// myHouse.Size = 28
	// myHouse.Price = 10
	// myHouse.Category = "아파트"
	var myHouse House = House{Address: "서울시 구로구"}

	fmt.Printf("주소: %s, 평수: %d", myHouse.Address, myHouse.Size)
}
