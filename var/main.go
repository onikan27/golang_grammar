package main

import "fmt"

func main() {
	var i int = 1
	var f64 float64 = 1.2
	// 同一の型なら下記のように定義できる
	var t, f bool = true, false
	fmt.Println(i, f64, t, f)
	// 1 1.2 true false

	// 下記のように複数の変数を宣言することもできる。
	var (
		hoge string = "hoge"
		foo  string = "foo"
	)

	fmt.Println(hoge, foo)
	// hoge foo

	// 変数宣言した瞬間にゼロ値が格納される
	var (
		nullInt  int
		nullF64  float64
		nullStr  string
		nullBool bool
	)

	fmt.Println(nullInt, nullF64, nullStr, nullBool)
	// => 0 0  false

	// varを使わなくても変数に値を代入できる。
	// 条件1：関数内であること（varは関数外でも宣言可能）
	str := "文字列"
	fmt.Println(str)

	xf64 := 1.2
	fmt.Printf("%T\n", xf64)
	// float64

	// 型定義を明確にしたい（デフォ値出ない挙動にしたい場合はvarを使う）
	var xf32 float32 = 1.2
	fmt.Printf("%T\n", xf32)
	// float32
}
