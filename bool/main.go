package main

import "fmt"

func main() {
	t, f := true, false
	fmt.Printf("type=%T value=%v", t, t)
	fmt.Printf("type=%T value=%v", f, f)

	// %tで厳密にbool型でないと出力されないようにできる
	fmt.Printf("%t", t)
	// !t(int=1)と出力される
	fmt.Printf("%t", 1)

	fmt.Println(!true)
	fmt.Println(!false)
}
