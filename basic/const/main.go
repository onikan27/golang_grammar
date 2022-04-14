package main

import "fmt"

func main() {
	const hello string = "hello world"
	fmt.Println(hello)

	// 定数の場合は右辺を省略できる
	const (
		a = 1 + 2
		b
		c
	)
	fmt.Println(a, b, c)
}
