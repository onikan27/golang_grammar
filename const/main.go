package main

import "fmt"

// constで定数を宣言することが可能
const Pi = 3.14

// varと同様に複数宣言が可能
const (
	UserName = "testuser"
	Password = "testpassword"
)

func main() {
	fmt.Println(Pi, UserName, Password)
}
