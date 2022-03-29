package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello" + "World")

	// 文字列の最初を表示したい場合はstringでキャストする
	fmt.Println(string("Hello"[0]))

	var s string = "Hello"

	// 最初のHをXに変換する
	// 実際のsは変わっていない（非破壊的）
	fmt.Println(strings.Replace(s, "H", "X", 1))

	// 特定の文字列が存在するか確認する
	fmt.Println(strings.Contains(s, "Hello"))

	fmt.Println("\"")
	// ``で囲ってもok
	fmt.Println(`"`)
}
