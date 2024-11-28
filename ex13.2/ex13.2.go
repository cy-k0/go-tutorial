package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"송하나", "hana", 25}
	vip := VIPUser{
		User{"송둘째", "brother", 20},
		3,
		250,
	}
	fmt.Printf("유저: %s, ID: %s, 나이: %d\n", user.Name, user.ID, user.Age)
}
