package main

// Golangの標準パッケージ
import "fmt"

// 一番初めに実行される関数
// 初期設定で使われる
func init() {
	fmt.Println("init")
}

// main関数がないとエラーになる
func main() {
	fmt.Println("Hello")
	hoge()
}

func hoge() {
	fmt.Println("hoge")
	fmt.Println("hoge", "foo")
}
