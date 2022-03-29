package main

import "fmt"

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	// Golangの配列はサイズを変更できない
	// スライスはサイズが変更可能
	// 宣言と同時に値を格納する場合は下記のように記述する
	var b [2]int = [2]int{10, 20}
	fmt.Println(b)

	var c []int = []int{100, 200}
	// スライスは可変長
	c = append(c, 300)
	fmt.Println(c)
}
