package main

// 複数のパッケージをimportするには()で囲む
import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("hello", time.Now())
	fmt.Println(user.Current())
}
