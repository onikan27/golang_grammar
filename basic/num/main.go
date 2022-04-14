package main

import "fmt"

func main() {
	// 数値はアンスコで区切っても問題ない
	const num = 100_000
	fmt.Println(num)
	// output:100000
}
