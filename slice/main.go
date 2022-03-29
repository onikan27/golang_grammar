package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5}
	// n[1:3]のように書くことで[2 3]が返ってくる
	// カンマで数えるとわかりやすいかも
	fmt.Println(n[1:3])
	fmt.Println(n[2:3])

	// 例えば最初の三つを抜き出したい場合は下記のように書く
	fmt.Println(n[0:3])
	// 0は省略できる
	fmt.Println(n[:3])
	// 最後のindexも省略できる
	fmt.Println(n[3:])

	// 最後から3番目なら下記のように記述可能
	fmt.Println(n[len(n)-3:])

	// sliceの中にsliceを格納することも可能
	var board = [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
	}

	fmt.Println(board)
}
